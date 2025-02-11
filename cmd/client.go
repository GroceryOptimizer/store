package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GroceryOptimizer/store/errors"
	grocer "github.com/GroceryOptimizer/store/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientHandshake(ctx context.Context, config string) (*grpc.ClientConn, *grocer.HandShakeResponse, error) {
	addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if addr == "" {
		addr = "localhost:5241" // fallback if not set
	}

	storeName := os.Getenv("STORE_NAME")
	if storeName == "" {
		errors.ErrStoreNameEnv("STORE_NAME env is not set")
	}

	store := grocer.Store{Name: storeName}

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
		conn, err := grpc.Dial(addr,
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
		storeId, err := client.HandShake(ctx, &grocer.HandShakeRequest{Store: &store})
		cancel()
		conn.Close()
		if err != nil {
			lastErr = errors.ErrClientHandshake("Failed to handshake with gRPC server", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Store ID:", storeId)
		return conn, storeId, nil
	}

}
