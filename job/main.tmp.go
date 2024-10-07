package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const apiKey = "cs1m7qpr01qsperucqbgcs1m7qpr01qsperucqc0" // Replace with your Finnhub API key

type StockQuote struct {
	CurrentPrice  float64 `json:"c"`
	HighPrice     float64 `json:"h"`
	LowPrice      float64 `json:"l"`
	OpenPrice     float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
	Volume        float64 `json:"v"`
}

func getStockQuote(symbol string) (*StockQuote, error) {
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	var quote StockQuote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		return nil, err
	}

	return &quote, nil
}

func main() {
	symbol := "TIGR" // Example: Apple Inc.

	stockQuote, err := getStockQuote(symbol)
	if err != nil {
		fmt.Println("Error fetching stock quote:", err)
		os.Exit(1)
	}

	fmt.Printf("Current Price: %.2f\n", stockQuote.CurrentPrice)
	fmt.Printf("High Price: %.2f\n", stockQuote.HighPrice)
	fmt.Printf("Low Price: %.2f\n", stockQuote.LowPrice)
	fmt.Printf("Open Price: %.2f\n", stockQuote.OpenPrice)
	fmt.Printf("Previous Close Price: %.2f\n", stockQuote.PreviousClose)
	fmt.Printf("Volume: %.2f\n", stockQuote.Volume)
}
