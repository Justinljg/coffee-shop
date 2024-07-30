package coffee

import (
	"sync"
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/coffee"
)

func TestSimulateCustomerArrivals(t *testing.T) {
	customers := make(chan coffee.Customer, 100)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go coffee.SimulateCustomerArrivals(customers, &wg)

	time.Sleep(5 * time.Second) // Allow some time for customers to be added

	close(customers)
	customerCount := 0
	for range customers {
		customerCount++
	}

	if customerCount == 0 {
		t.Error("Expected some customers, got 0")
	}
}
