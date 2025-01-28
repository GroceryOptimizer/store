package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	vendor "github.com/GroceryOptimizer/store/proto"
)

type server struct {
	vendor.UnimplementedVendorServiceServer
}

func (s *server) SendMessage(ctx context.Context, req *vendor.SendMessageRequest) (*vendor.SendMessageReply, error) {
	fmt.Printf("Recieved from client: %s\n", req.GetMessage())
	resp := &vendor.SendMessageReply{
		Reply: "Hello from Go server, you said: " + req.GetMessage(),
	}
	return resp, nil
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	vendor.RegisterVendorServiceServer(grpcServer, &server{})

	log.Println("gRPC Go server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}


	//http.ListenAndServe(":8080", nil)
}
