package cafe

import (
	"context"
	"math/rand"
	"time"
	"fmt"
)

// Customer represents a customer with a unique ID.
type Customer struct {
	ID int
}

// SimulateCustomerArrivals continuously generates customer arrivals and sends them to the customers channel.
// This function runs until the context is cancelled or the parent goroutine signals completion by closing the channel.
// It increments the customer ID for each new customer.
func SimulateCustomerArrivals(ctx context.Context, customers chan<- Customer, numCustomers int) {
	for i := 1; i <= numCustomers; i++ {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(rand.Intn(3)) * time.Second):
			fmt.Printf("Customer %d arrives.\n", i)
			// Send a new customer to the channel
			customers <- Customer{ID: i}
		}
	}
}
