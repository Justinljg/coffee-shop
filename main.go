package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/justinljg/coffee-shop/coffee"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	customers := make(chan coffee.Customer, 10)
	orders := make(chan coffee.Order, 10)
	baristas := []coffee.Barista{{Id: 1}, {Id: 2}}
	var wg sync.WaitGroup

	// Handle signals to allow graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Simulate customer arrivals
	wg.Add(1)
	go coffee.SimulateCustomerArrivals(customers, &wg)

	// Handle orders
	for _, barista := range baristas {
		wg.Add(1)
		go func(b coffee.Barista) {
			for customer := range customers {
				order := coffee.Order{CustomerID: customer.Id, CoffeeType: coffee.CoffeeType(rng.Intn(3))}
				fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.Id, coffee.CoffeeTypeToString(order.CoffeeType))
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
		fmt.Printf("Customer %d receives their %s and leaves.\n", order.CustomerID, coffee.CoffeeTypeToString(order.CoffeeType))
	}

	fmt.Println("Coffee shop closed.")
}
