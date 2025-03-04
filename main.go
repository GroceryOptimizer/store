package main

import (
	"context"
	"os"
	"sync"
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
	log.Printf("port: %v", port)
	if port == "" {
		port = "12345"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen to %s: %v", port, err)
	}
	log.Printf("gRPC Go server listening on port %v", port)
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
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn := &grpc.ClientConn{}
	res := &grocer.HandShakeResponse{}
	wg.Add(1)
	go func () {
		defer wg.Done()
		conn, res, _ = cmd.ClientHandshake(ctx, serviceConfig)
	}()
	wg.Wait()
	defer conn.Close()

	resp, _ := cmd.UpdateInventory(ctx, conn, res.Id)
	log.Println("inventory response: ", resp)

	grpcServer := grpc.NewServer()

	grocer.RegisterStoreServiceServer(grpcServer, &cmd.Server{})

	//fmt.Println(os.Getenv("STORE_NAME"))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
