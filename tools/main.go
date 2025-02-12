package tools

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/GroceryOptimizer/store/errors"
	"github.com/GroceryOptimizer/store/proto"
)

func GetClientAddress() string {
	var host_port []string
	var store_addr string
	host := os.Getenv("STORE_HOST")
	if ips, err := net.LookupHost(host); err == nil && len(ips) > 0 {
		host = ips[0]
		host_port = append(host_port, host)
	}
	port := os.Getenv("STORE_PORT")
	if p, err := net.LookupPort("tcp", port); err == nil {
		host_port = append(host_port, strconv.Itoa(p))
		store_addr = strings.Join(host_port, ":")
	}
	return store_addr
}

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

