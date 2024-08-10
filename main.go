package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func main() {
	var numCustomers int
	flag.IntVar(&numCustomers, "numCustomers", 100, "Number of customers to simulate")
	flag.Parse()

	// Make the channels
	customers := make(chan cafe.Customer, 10)
	orders := make(chan cafe.Order, 10)
	baristas := []cafe.Barista{{ID: 1}, {ID: 2}}

	// Create waitgroup
	var wg sync.WaitGroup

	// Set up context and cancellation function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate customer arrivals
	wg.Add(1)
	go func() {
		defer wg.Done()
		cafe.SimulateCustomerArrivals(ctx, customers, numCustomers)
		// Close the customers channel after all customers are sent
		close(customers)
	}()

	// Start baristas with cancellation context
	for _, barista := range baristas {
		wg.Add(1)
		rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(barista.ID)))
		go barista.ServeCustomers(ctx, customers, orders, &wg, rng)
	}

	// Goroutine to handle the signal and initiate shutdown
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		fmt.Println("\nInterrupt signal received. Shutting down...")
		cancel()
	}()

	// Close the orders channel once all orders are prepared
	go func() {
		wg.Wait()
		close(orders)
	}()

	// Customer collects their orders from the orders channel
	for order := range orders {
		fmt.Printf("Customer %d receives their %s and leaves.\n", order.CustomerID, cafe.CoffeeTypeToString(order.CoffeeType))
	}

	wg.Wait()
	fmt.Println("Coffee shop closed.")
}
