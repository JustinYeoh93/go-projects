package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

type MarketData struct {
	Exchange  string  `json:"exchange"`
	Symbol    string  `json:"symbol"`
	LastPrice float64 `json:"last_price"`
	Bid       float64 `json:"bid"`
	Ask       float64 `json:"ask"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Timestamp string  `json:"timestamp"`
}

func main() {
	xetraMart := `{"exchange": "Xetra", "symbol": "AAPL", "last_price": 142.50, "change": -1.20, "percent_change": -0.83, "volume": 25000, "bid": 142.45, "ask": 142.55, "high": 143.10, "low": 142.30, "timestamp": "2023-06-24T09:35:12Z", "market_cap": 200000000000, "pe_ratio": 25.5}`
	tradegateMart := `{"exchange": "Tradegate", "symbol": "AAPL", "last_price": 142.55, "cur_change": -1.20, "per_change": -0.83, "volume": 18000, "bid": 142.50, "ask": 142.60, "high": 143.10, "low": 142.30, "timestamp": "2023-06-24T09:35:12Z", "dividend_yield": 1.5}`
	comdirectMart := `{"exchange": "Comdirect", "symbol": "AAPL", "last_price": 142.60, "change": -1.20, "percent_change": -0.83, "volume": 20000, "bid": 142.55, "ask": 142.65, "high": 143.10, "low": 142.30, "timestamp": "2023-06-24T09:35:12Z"}`

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Process Xetra market data concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		processMart(xetraMart)
	}()

	// Process Tradegate market data concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		processMart(tradegateMart)
	}()

	// Process Comdirect market data concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		processMart(comdirectMart)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

func processMart(jsonData string) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create an instance of the target struct
	MarketData := MarketData{}

	// Iterate over the fields of the struct using reflection
	userType := reflect.TypeOf(MarketData)
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fieldValue, exists := data[field.Tag.Get("json")]
		if exists {
			reflectValue := reflect.ValueOf(&MarketData).Elem().Field(i)

			// Handle type conversions dynamically based on the field type
			switch reflectValue.Type().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if reflectValue.CanSet() {
					intValue, ok := fieldValue.(int)
					if ok {
						reflectValue.SetInt(int64(intValue))
					}
				}
			case reflect.Float32, reflect.Float64:
				if reflectValue.CanSet() {
					floatValue, ok := fieldValue.(float64)
					if ok {
						reflectValue.SetFloat(float64(floatValue))
					}
				}
			case reflect.String:
				if reflectValue.CanSet() {
					strValue, ok := fieldValue.(string)
					if ok {
						reflectValue.SetString(strValue)
					}
				}
			}
		}
	}

	fmt.Printf("Exchange=%s, MarketData:\n%+v\n", MarketData.Exchange, MarketData)
}
