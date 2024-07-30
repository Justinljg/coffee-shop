package main

import (
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
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	customers := make(chan cafe.Customer, 10)
	orders := make(chan cafe.Order, 10)
	baristas := []cafe.Barista{{ID: 1}, {ID: 2}}
	var wg sync.WaitGroup

	// Handle signals to allow graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Simulate customer arrivals
	wg.Add(1)
	go cafe.SimulateCustomerArrivals(customers, &wg)

	// Handle orders
	for _, barista := range baristas {
		wg.Add(1)
		go func(b cafe.Barista) {
			for customer := range customers {
				order := cafe.Order{CustomerID: customer.ID, CoffeeType: cafe.CoffeeType(rng.Intn(3))}
				fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.ID, cafe.CoffeeTypeToString(order.CoffeeType))
				wg.Add(1)
				go b.PrepareOrder(order, orders, &wg)
			}
			wg.Done()
		}(barista)
	}

	// Wait for stop signal
	<-stop

	// Cleanup
	close(customers)
	wg.Wait()
	close(orders)

	for order := range orders {
		fmt.Printf("Customer %d receives their %s and leaves.\n", order.CustomerID, cafe.CoffeeTypeToString(order.CoffeeType))
	}

	fmt.Println("Coffee shop closed.")
}
