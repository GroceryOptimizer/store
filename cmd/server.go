package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"github.com/GroceryOptimizer/store/proto"
	"github.com/GroceryOptimizer/store/tools"
)

type Server struct {
	grocer.UnimplementedStoreServiceServer
}

// gRPC method to retrieve product prices
func (s *Server) Products(ctx context.Context, req *grocer.InventoryRequest) (*grocer.InventoryResponse, error) {
	fmt.Println("Received ShoppingCart message:", len(req.GetShoppingCart()), "products")

	// Read stock items from JSON file
	stockItems, err := tools.ReadJSONFile("./products.json")
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

// gRPC method to capitalize a message
func (s *Server) SendMessage(ctx context.Context, req *grocer.SendMessageRequest) (*grocer.SendMessageResponse, error) {
	originalMsg := req.GetMessage()
	capitalizedMsg := strings.ToUpper(originalMsg)
	fmt.Printf("Received from client: %s\n", req.GetMessage())

	resp := &grocer.SendMessageResponse{
		Reply: "Hello from Go server, you said: " + capitalizedMsg,
	}
	return resp, nil
}

