package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GroceryOptimizer/store/errors"
	grocer "github.com/GroceryOptimizer/store/proto"
	"github.com/GroceryOptimizer/store/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func ClientHandshake(ctx context.Context, config string) (*grpc.ClientConn, *grocer.HandShakeResponse, error) {
	hub_addr := os.Getenv("GRPC_SERVER_ADDRESS")
	log.Println("Hub Address: ", hub_addr)
	if hub_addr == "" {
		hub_addr = "localhost:5241" // fallback if not set
	}

	storeName := os.Getenv("STORE_NAME")
	if storeName == "" {
		errors.ErrStoreNameEnv("STORE_NAME env is not set")
	}

	store_addr := tools.GetClientAddress()

	loc := tools.GetStoreCoords()

	store := grocer.Store{Name: storeName, GrpcAddress: store_addr, Location: &loc}

	deadline, ok := ctx.Deadline()
	if !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		deadline, _ = ctx.Deadline()
	}
	var lastErr error

	for {
		if time.Now().After(deadline) {
			log.Printf("handshake timed out after 30 seconds")
			return nil, nil, lastErr
		}
		conn, err := grpc.Dial(hub_addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(config),
		)
		if err != nil {
			lastErr = errors.ErrClientHandshake("Failed to connect to gRPC server", err)
			time.Sleep(5 * time.Second)
			continue
		}

		client := grocer.NewHubServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		res, err := client.HandShake(ctx, &grocer.HandShakeRequest{Store: &store})
		cancel()
		//conn.Close()
		if err != nil {
			lastErr = errors.ErrClientHandshake("Failed to handshake with gRPC server", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Store ID:", res)
		return conn, res, nil
	}

}

func SendInventoryList(ctx context.Context, conn *grpc.ClientConn, storeId string) (*grocer.UpdateInventoryResponse, error) {
	//defer conn.Close()
	log.Println("conn: ", conn)
	log.Println("storeId: ", storeId)

	client := grocer.NewHubServiceClient(conn)

	stockItems, err := tools.ReadJSONFile("./products.json")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read json file: %v", err)
	}

	// Send update request to Hub
	hubReq := &grocer.UpdateInventoryRequest{
		StoreId:   storeId,
		StockItems: stockItems,
	}

	hubResp, err := client.UpdateInventory(ctx, hubReq)
	if err != nil {
		log.Println("Failed to update inventory in Hub: ", err)
		return nil, status.Errorf(codes.Internal, "Failed to update inventory in Hub: %v", err)
	}

	// Return response to Store
	return &grocer.UpdateInventoryResponse{
		Message: hubResp.Message,
	}, nil
}
