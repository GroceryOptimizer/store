package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GroceryOptimizer/store/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientHandshake(ctx context.Context, config string) *grocer.HandShakeResponse {
	addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if addr == "" {
		addr = "localhost:5241" // fallback if not set
	}

	storeName := os.Getenv("STORE_NAME")
	if storeName == "" {
		log.Fatalf("STORE_NAME environment variable not set")
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
		if time.Now() == deadline {
			log.Fatalf("Handshake timed out: %v", lastErr)
		}
		conn, err := grpc.Dial(addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(config),
			)
		if err != nil {
			lastErr = err
			log.Printf("Failed to connect to gRPC server: %v", err)
			time.Sleep(5)
			continue
		}

		client := grocer.NewHubServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		storeId, err := client.HandShake(ctx, &grocer.HandShakeRequest{Store: &store})
		cancel()
		conn.Close()
		if err != nil {
			lastErr = err
			log.Printf("Failed to handshake with gRPC server: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Store ID:", storeId)
		return storeId
	}

}
