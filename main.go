package main

import (
	"context"
	"strings"
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
	originalMsg := req.GetMessage()
	capitalizedMsg := strings.ToUpper(originalMsg)
	fmt.Printf("Recieved from client: %s\n", req.GetMessage())
	resp := &vendor.SendMessageReply{
		Reply: "Hello from Go server, you said: " + capitalizedMsg,
	}
	return resp, nil
}

func (s *server) Products(ctx context.Context, req *vendor.InventoryRequest) (*vendor.InventoryReply, error){
	fmt.Println("Recieved ShoppingCart message: ", len(req.GetShoppingCart()), "products")

	//################################### LOGIC WITH STOCKITEMS GOES HERE ##########################################
	// For illustration, create StockItems with some dummy prices
    // or logic that sets real prices:
    var items []*vendor.StockItem
    for _, p := range req.GetShoppingCart() {
        items = append(items, &vendor.StockItem{
            Name:  p.GetName(),
            Price: 42,  // or compute a real price
        })
    }
	//##############################################################################################################

	resp := &vendor.InventoryReply{
		StockItems: items,
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
