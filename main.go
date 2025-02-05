package main

import (
	"context"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"

	store "github.com/GroceryOptimizer/store/proto"
)

type server struct {
	store.UnimplementedStoreServiceServer
}

// gRPC method to capitalize a message
func (s *server) SendMessage(ctx context.Context, req *store.SendMessageRequest) (*store.SendMessageResponse, error) {
	originalMsg := req.GetMessage()
	capitalizedMsg := strings.ToUpper(originalMsg)
	fmt.Printf("Received from client: %s\n", req.GetMessage())

	resp := &store.SendMessageResponse{
		Reply: "Hello from Go server, you said: " + capitalizedMsg,
	}
	return resp, nil
}

// gRPC method to retrieve product prices
func (s *server) Products(ctx context.Context, req *store.InventoryRequest) (*store.InventoryResponse, error) {
	fmt.Println("Received ShoppingCart message:", len(req.GetShoppingCart()), "products")

	// Read stock items from JSON file
	stockItems, err := readJSONFile("./products.json")
	if err != nil {
		log.Fatalf("Failed to read the JSON file: %v", err)
	}

	// Convert stock list into a map for fast lookup
	stockMap := make(map[string]int32)
	for _, item := range stockItems {
		stockMap[item.Name] = item.Price
	}

	// Filter stock items based on requested shopping cart
	var items []*store.StockItem
	for _, p := range req.GetShoppingCart() {
		var name = strings.ToLower(p.Name)
		if price, found := stockMap[name]; found {
			items = append(items, &store.StockItem{
				Name:  name,
				Price: price,
			})
		} else {
			log.Printf("Product not found in stock: %s", name)
		}
	}

	resp := &store.InventoryResponse{
		StockItems: items,
	}
	return resp, nil
}

// Read JSON file directly into a slice of gRPC StockItem messages
func readJSONFile(filename string) ([]*store.StockItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Use a generic map to parse JSON without custom structs
	var jsonData map[string][]map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		return nil, err
	}

	// Extract "stock" key and convert into []*store.StockItem
	var stockItems []*store.StockItem
	for _, item := range jsonData["stock"] {
		if product, ok := item["product"].(map[string]interface{}); ok {
			if name, exists := product["name"].(string); exists {
				if price, exists := item["price"].(float64); exists { // JSON numbers are float64 by default
					stockItems = append(stockItems, &store.StockItem{
						Name:  name,
						Price: int32(price),
					})
				}
			}
		}
	}

	return stockItems, nil
}

// gRPC Server Initialization
func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {

		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	store.RegisterStoreServiceServer(grpcServer, &server{})
	fmt.Println(os.Getenv("STORE_NAME"))

	log.Println("gRPC Go server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
