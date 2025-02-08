package main

import (
	"context"
	"time"

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

func clientHandshake(config string) *grocer.HandShakeResponse {
	addr := os.Getenv("GRPC_SERVER_ADDRESS")
	if addr == "" {
		addr = "localhost:5241" // fallback if not set
	}

	storeName := os.Getenv("STORE_NAME")
	if storeName == "" {
		log.Fatalf("STORE_NAME environment variable not set")
	}

	store := grocer.Store{Name: storeName}

	deadline := time.Now().Add(30 * time.Second)
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

// gRPC Server Initialization
func main() {
	lis, err := net.Listen("tcp", ":50051")
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

	clientHandshake(serviceConfig)

	grpcServer := grpc.NewServer()

	grocer.RegisterStoreServiceServer(grpcServer, &server{})
	//fmt.Println(os.Getenv("STORE_NAME"))

	log.Println("gRPC Go server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
