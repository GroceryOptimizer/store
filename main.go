package main

import (
	"context"
	"os"
	"time"

	"log"
	"net"

	"github.com/GroceryOptimizer/store/cmd"
	"google.golang.org/grpc"

	grocer "github.com/GroceryOptimizer/store/proto"
)

// gRPC Server Initialization
func main() {
	port := os.Getenv("STORE_PORT")
	if len(port) == 0 {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	serviceConfig := `{
	  "methodConfig": [{
	    "name": [{"service": "grocer.StoreService"}],
	    "retryPolicy": {
	      "maxAttempts": 5,
	      "initialBackoff": "1s",
	      "maxBackoff": "5s",
	      "backoffMultiplier": 1.5,
	      "retryableStatusCodes": ["UNAVAILABLE"]
	    }
	  }]
	}`
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cmd.ClientHandshake(ctx, serviceConfig)
	defer cancel()

	grpcServer := grpc.NewServer()

	grocer.RegisterStoreServiceServer(grpcServer, &cmd.Server{})

	//fmt.Println(os.Getenv("STORE_NAME"))

	log.Printf("gRPC Go server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
