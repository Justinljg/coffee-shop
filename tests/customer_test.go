package cafe

import (
	"testing"
	"time"
	"github.com/justinljg/coffee-shop/cafe"
)

func TestSimulateCustomerArrivals(t *testing.T) {
	numCustomers := 10
	customers := make(chan cafe.Customer, numCustomers)

	// Start SimulateCustomerArrivals in a separate goroutine
	go cafe.SimulateCustomerArrivals(customers, numCustomers)

	// Wait for all customers to arrive and check if the correct number of customers was received
	for i := 1; i <= numCustomers; i++ {
		select {
		case customer := <-customers:
			if customer.ID != i {
				t.Errorf("Expected customer ID %d, but got %d", i, customer.ID)
			}
		case <-time.After(5 * time.Second): // Adjust timeout as needed
			t.Error("Customer did not arrive in the channel within the expected time")
		}
	}

	// Ensure no more customers are sent to the channel
	select {
	case _, open := <-customers:
		if open {
			t.Error("Expected no more customers, but received one")
		}
	default:
		// Channel is closed or empty, which is expected
	}
}
