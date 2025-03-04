package tools

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/GroceryOptimizer/store/errors"
	grocer "github.com/GroceryOptimizer/store/proto"
)

func GetStoreCoords() grocer.Coordinates {
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

type ProductJSON struct {
	Name string `json:"name"`
}

type StockItemJSON struct {
	Product  ProductJSON `json:"product"`
	Quantity int32       `json:"quantity"`
	Price    int32       `json:"price"`
}

type StockData struct {
	Stock []StockItemJSON `json:"stock"`
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
		log.Printf("Error reading file: %v", err)
		return nil, errors.ErrDatabaseFailure(err)
	}

	// Use a generic map to parse JSON without custom structs
	var jsonData StockData
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		return nil, err
	}

	// TODO: // This may be broken, check it
	// Extract "stock" key and convert into []*grocer.StockItem
	var stockItems []*grocer.StockItem
	for _, item := range jsonData.Stock {
	 stockItem := &grocer.StockItem{
				Product: &grocer.Product{Name: item.Product.Name},
				Quantity: item.Quantity,
				Price: item.Price,
		}
		log.Printf("StockItem: %v", stockItem)
		stockItems = append(stockItems, stockItem)
	}

	return stockItems, nil
}
