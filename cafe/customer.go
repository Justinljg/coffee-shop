package cafe

import (
	"math/rand"
	"sync"
	"time"
)

// Customer represents a customer with a unique ID.
type Customer struct {
	ID int
}

// SimulateCustomerArrivals continuously generates customer arrivals and sends them to the customers channel.
// This function runs until the parent goroutine signals completion by closing the channel.
// It increments the customer ID for each new customer.
func SimulateCustomerArrivals(customers chan<- Customer, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when the function returns

	customerID := 1 // Initial customer ID

	// Infinite loop to simulate continuous customer arrivals
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Random sleep to simulate random arrival times

		// Send a new customer with a unique ID to the customers channel
		customers <- Customer{ID: customerID}
		customerID++ // Increment customer ID for the next customer
	}
}
