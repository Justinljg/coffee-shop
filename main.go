package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func main() {
	// make the channels
	customers := make(chan cafe.Customer, 10)
	orders := make(chan cafe.Order, 10)
	baristas := []cafe.Barista{{ID: 1}, {ID: 2}}

	// create waitgroup
	var wg sync.WaitGroup

	// set number of customers to 100
	numCustomers := 100

	// Simulate customer arrivals
	wg.Add(1)
	go func() {
		defer wg.Done()
		cafe.SimulateCustomerArrivals(customers, numCustomers)
		close(customers)
	}()

	// Barista to prepare orders for loop as two baristas
	for _, barista := range baristas {
		wg.Add(1)
		go func(b cafe.Barista) {
			defer wg.Done()
			// Each barista has its own rand.Rand instance to prevent data race.
			rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(b.ID)))
			for customer := range customers {
				order := cafe.Order{CustomerID: customer.ID, CoffeeType: cafe.CoffeeType(rng.Intn(3))}
				fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.ID, cafe.CoffeeTypeToString(order.CoffeeType))				
				// prepare order and send completed orders to orders channel
				b.PrepareOrder(order, orders)
			}
		}(barista)
	}

	// Close the orders channel once all orders are prepared
	go func() {
		wg.Wait()
		close(orders)
	}()

	// Customer collects their orders from the orders channel
	for order := range orders {
		fmt.Printf("Customer %d receives their %s and leaves.\n", order.CustomerID, cafe.CoffeeTypeToString(order.CoffeeType))
	}

	fmt.Println("Coffee shop closed.")
}
