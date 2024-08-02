package cafe

import (
	"math/rand"
	"time"
)

// Customer represents a customer with a unique ID.
type Customer struct {
	ID int
}

// SimulateCustomerArrivals continuously generates customer arrivals and sends them to the customers channel.
// This function runs until the parent goroutine signals completion by closing the channel.
// It increments the customer ID for each new customer.
func SimulateCustomerArrivals(customers chan<- Customer, numCustomers int) {
	for i := 1; i <= numCustomers; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		customers <- Customer{ID: i}
	}
}
