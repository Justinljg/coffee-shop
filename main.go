package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"


	"github.com/justinljg/coffee-shop/cafe"
)

func main() {
	var numCustomers int
	flag.IntVar(&numCustomers, "numCustomers", 100, "Number of customers to simulate")
	flag.Parse()

	customers := make(chan cafe.Customer, 10)
	orders := make(chan cafe.Order, 10)
	baristas := []cafe.Barista{{ID: 1}, {ID: 2}}

	var wg sync.WaitGroup

	// Create a context that supports cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called to free resources

	// Goroutine to simulate customer arrivals
	wg.Add(1)
	go func() {
		defer wg.Done()
		cafe.SimulateCustomerArrivals(ctx, customers, numCustomers)
		close(customers)
	}()

	for _, barista := range baristas {
		wg.Add(1)
		go func(b cafe.Barista) {
			defer wg.Done()
			b.ServeCustomers(ctx, customers, orders)
		}(barista)
	}	

	// Goroutine to close the orders channel once all orders are processed
	go func() {
		wg.Wait() // Wait for all barista goroutines to finish
		close(orders) // Close orders channel
	}()

	// Goroutine to handle interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nInterrupt signal received. Shutting down...")
		cancel() // Signal cancellation to all goroutines
	}()

	// Process orders
	for order := range orders {
		fmt.Printf("Customer %d receives their %s and leaves.\n", order.CustomerID, cafe.CoffeeTypeToString(order.CoffeeType))
	}

	// Ensure all remaining orders are processed before exiting
	wg.Wait()
	fmt.Println("Coffee shop closed.")
}
