package cafe

import (
	"sync"
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestSimulateCustomerArrivals(t *testing.T) {
	customers := make(chan cafe.Customer, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go cafe.SimulateCustomerArrivals(customers, &wg)

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
