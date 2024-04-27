package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

type Customer struct {
	Code int64
}

type Product struct {
	Code     int64
	Quantity int8
}

type Payment struct {
	Card         Card
	Billing      Billing
	Transaction  Transaction
	Verification Verification
}

type Card struct {
	Pan                   int64
	ExpirationDate        time.Time
	CardholderName        string
	CardVerificationToken string
}

type Billing struct {
	Address string
	ZipCode string
	Country string
}

type Transaction struct {
	Amount            float64
	CurrencyCode      string
	PaymentMethodCode string
	OrderNumber       int64
}

type Shipping struct {
	Geolocation GeoLocation
	Address     AddressCl
	Receiver    Receiver
}

type GeoLocation struct {
	Latitude  float64
	Longitude float64
}

type AddressCl struct {
	RegionCode  int8
	RegionName  string
	ComunaCode  int16
	ComunaName  string
	CalleName   string
	CalleNumber string
	Comments    string
}

type Receiver struct {
	Name string
}

type Verification struct {
	AuthorizationCode string
	TransactionStatus string
	TransactionId     string
	Timestamp         time.Time
}

type Sale struct {
	IdSale   int64
	Customer Customer
	Products []Product
	Payments []Payment
	Shipping Shipping
}

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("QUEUE_REDIS_IP") + ":" + os.Getenv("QUEUE_REDIS_PORT"),
		Password: os.Getenv("QUEUE_REDIS_PASSWORD"),
		DB:       0,
	})

	// Start worker to process sales
	go processSales(rdb)

	// Wait for termination signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	fmt.Println("Shutting down...")
}

func processSales(rdb *redis.Client) {
	for {
		// Pop the first sale from the queue
		saleJSON, err := rdb.LPop(context.Background(), "sales_queue").Result()
		if err != nil {
			fmt.Print(".") // fmt.Println("Failed to retrieve sale from the queue:", err)
			continue
		}
		fmt.Println(saleJSON)

		// Make HTTP POST request to the microservice
		resp, err := http.Post("http://localhost:2500/sale", "application/json", bytes.NewBuffer([]byte(saleJSON)))
		if err != nil {
			fmt.Printf("failed to make HTTP request: %v", err)
			continue
		}
		defer resp.Body.Close()

		// Log the response code
		fmt.Printf("Response Code: %d\n", resp.StatusCode)

		// Check response status code
		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("unexpected status code: %d", resp.StatusCode)
			continue
		}

		// // Check if response body is empty
		// if resp.ContentLength == 0 {
		// 	fmt.Println("Response body is empty")
		// 	continue
		// }

		// // Decode response body
		// var responseProduct Sale
		// if err := json.NewDecoder(resp.Body).Decode(&responseProduct); err != nil {
		// 	fmt.Printf("failed to decode response body: %v", err)
		// 	continue
		// }
	}
}
