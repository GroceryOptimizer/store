package tools

import (
	"encoding/json"
	"io"
	"github.com/GroceryOptimizer/store/proto"
	"os"
	"github.com/GroceryOptimizer/store/errors"
	
)

// Read JSON file directly into a slice of gRPC StockItem messages
func ReadJSONFile(filename string) ([]*grocer.StockItem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.ErrDatabaseFailure(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.ErrDatabaseFailure(err)
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

