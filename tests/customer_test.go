package cafe

import (
	"github.com/justinljg/coffee-shop/cafe"
	"testing"
)

func TestSimulateCustomerArrivals(t *testing.T) {
	numCustomers := 10
	customers := make(chan cafe.Customer, numCustomers)

	go func() {
		cafe.SimulateCustomerArrivals(customers, numCustomers)
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
