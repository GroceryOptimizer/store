package tools

import (
	"encoding/json"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"log"
	"github.com/GroceryOptimizer/store/errors"
	grocer "github.com/GroceryOptimizer/store/proto"
)

func GetStoreCoords() (grocer.Coordinates) {
	lat, err := strconv.ParseFloat(os.Getenv("LATITUDE"), 64)
	if err != nil {
		errors.ErrStoreNameEnv("LATITUDE env is not set")
	}

	long, err := strconv.ParseFloat(os.Getenv("LONGITUDE"), 64)
	if err != nil {
		errors.ErrStoreNameEnv("LONGITUDE env is not set")
	}

	return grocer.Coordinates{Latitude: lat, Longitude: long}
}

func GetClientAddress() string {
	var host_port []string
	var store_addr string

	hostname:= os.Getenv("STORE_NAME")
	log.Println("Hostname collected: ", hostname)

	host_port = append(host_port, hostname)

	port := os.Getenv("STORE_PORT")
	if p, err := net.LookupPort("tcp", port); err == nil{
		host_port = append(host_port, strconv.Itoa(p))
		store_addr = strings.Join(host_port, ":")
	}
	log.Println("About to send handshake with address: ", store_addr)
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
