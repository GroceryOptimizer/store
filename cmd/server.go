package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/GroceryOptimizer/store/errors"
	grocer "github.com/GroceryOptimizer/store/proto"
	"github.com/GroceryOptimizer/store/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	grocer.UnimplementedStoreServiceServer
}

// gRPC method to retrieve product prices
func (s *Server) Products(ctx context.Context, req *grocer.InventoryRequest) (*grocer.InventoryResponse, error) {
	if len(req.GetShoppingCart()) == 0 {
		return nil, errors.ErrInvalidRequest("Shopping cart is empty")
	}

	fmt.Println("Received ShoppingCart message:", len(req.GetShoppingCart()), "products")

	// Read stock items from JSON file
	stockItems, err := tools.ReadJSONFile("./products.json")
	if err != nil {
		return nil, errors.ErrDatabaseFailure(err)
	}

	if len(stockItems) == 0 {
		return nil, status.Error(codes.NotFound, "No products available in stock")
	}

	// Convert stock list into a map for fast lookup
	stockMap := make(map[string]int32)
	for _, item := range stockItems {
		stockMap[item.Product.Name]= item.Price
	}

	// Filter stock items based on requested shopping cart
	var items []*grocer.StockItem
	for _, p := range req.GetShoppingCart() {
		var name = strings.ToLower(p.Name)
		if price, found := stockMap[name]; found {
			items = append(items, &grocer.StockItem{
				Product:  &grocer.Product{Name: name},
				Price: price,
			})
		} else {
			errors.ErrProductNotFound(name)
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
