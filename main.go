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
	"google.golang.org/grpc/credentials/insecure"

	"github.com/GroceryOptimizer/store/proto"
)

type server struct {
	grocer.UnimplementedStoreServiceServer
}

// gRPC method to capitalize a message
func (s *server) SendMessage(ctx context.Context, req *grocer.SendMessageRequest) (*grocer.SendMessageResponse, error) {
	originalMsg := req.GetMessage()
	capitalizedMsg := strings.ToUpper(originalMsg)
	fmt.Printf("Received from client: %s\n", req.GetMessage())

	resp := &grocer.SendMessageResponse{
		Reply: "Hello from Go server, you said: " + capitalizedMsg,
	}
	return resp, nil
}

func NewContext(ctx context.Context, store *grocer.Store) context.Context {
	var storeKey = "STORE_NAME"
	return context.WithValue(ctx, storeKey, store)
}

// gRPC method to retrieve product prices
func (s *server) Products(ctx context.Context, req *grocer.InventoryRequest) (*grocer.InventoryResponse, error) {
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
	var items []*grocer.StockItem
	for _, p := range req.GetShoppingCart() {
		var name = strings.ToLower(p.Name)
		if price, found := stockMap[name]; found {
			items = append(items, &grocer.StockItem{
				Name:  name,
				Price: price,
			})
		} else {
			log.Printf("Product not found in stock: %s", name)
		}
	}

	resp := &grocer.InventoryResponse{
		StockItems: items,
	}
	return resp, nil
}

// Read JSON file directly into a slice of gRPC StockItem messages
func readJSONFile(filename string) ([]*grocer.StockItem, error) {
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

	// Extract "stock" key and convert into []*grocer.StockItem
	var stockItems []*grocer.StockItem
	for _, item := range jsonData["stock"] {
		if product, ok := item["product"].(map[string]interface{}); ok {
			if name, exists := product["name"].(string); exists {
				if price, exists := item["price"].(float64); exists { // JSON numbers are float64 by default
					stockItems = append(stockItems, &grocer.StockItem{
						Name:  name,
						Price: int32(price),
					})
				}
			}
		}
	}

	return stockItems, nil
}

func clientHandshake() {
	addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if addr == "" {
        addr = "localhost:5241" // fallback if not set
    }
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	storeName := os.Getenv("STORE_NAME")
	if storeName == "" {
		log.Fatalf("STORE_NAME environment variable not set")
	}
	store := grocer.Store{Name: storeName}

	client := grocer.NewHubServiceClient(conn)
	storeId, err := client.HandShake(context.Background(), &grocer.HandShakeRequest{Store: &store})
	if err != nil {
		log.Fatalf("Failed to handshake with gRPC server: %v", err)
	}
	fmt.Println("Store ID:", storeId)
}

// gRPC Server Initialization
func main() {
	lis, err := net.Listen("tcp", ":50051")
	//var opts *[]grpc.DialOption

	clientHandshake()
	if err != nil {

		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	grocer.RegisterStoreServiceServer(grpcServer, &server{})
	//fmt.Println(os.Getenv("STORE_NAME"))

	log.Println("gRPC Go server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
