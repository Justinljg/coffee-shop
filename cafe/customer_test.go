package cafe

import (
	"context"
	"testing"
)

func TestSimulateCustomerArrivals(t *testing.T) {
	numCustomers := 10
	customers := make(chan Customer, numCustomers)

	// Set up context and cancellation function
	ctx := context.Background()

	go func() {
		SimulateCustomerArrivals(ctx, customers, numCustomers)
		close(customers)
	}()

	received := 0
	for customer := range customers {
		received++
		if customer.ID != received {
			t.Errorf("Expected customer ID %d, but got %d", received, customer.ID)
		}
	}

	if received != numCustomers {
		t.Errorf("Expected %d customers, but received %d", numCustomers, received)
	}
}
