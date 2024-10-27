package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const apiKey = "cs1m7qpr01qsperucqbgcs1m7qpr01qsperucqc0"

type StockQuote struct {
	CurrentPrice  float64 `json:"c"`
	HighPrice     float64 `json:"h"`
	LowPrice      float64 `json:"l"`
	OpenPrice     float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
	Volume        float64 `json:"v"`
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second, // Set timeout to 5 seconds
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/proxy/stock/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		// https://finnhub.io/api/v1/quote?symbol=tiger&token=cs1m7qpr01qsperucqbgcs1m7qpr01qsperucqc0
		url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", name, apiKey)

		// Make an HTTP request to the new URL
		resp, err := client.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			return
		}

		// Set the content type and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome!!!"))
	})
	http.ListenAndServe(":3000", r)
}
