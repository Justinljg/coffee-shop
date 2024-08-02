package cafe

import (
	"testing"
	"time"
	"sync"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestPrepareOrder(t *testing.T) {
	// Create a barista and an order
	barista := cafe.Barista{ID: 1}
	order := cafe.Order{
		CustomerID: 1,
		CoffeeType: cafe.Espresso, // Use a valid CoffeeType constant
	}

	// Create a channel to receive the order
	orders := make(chan cafe.Order)

	// Use a WaitGroup to wait for the completion of the PrepareOrder function
	var wg sync.WaitGroup
	wg.Add(1)

	// Run PrepareOrder in a separate goroutine
	go func() {
		defer wg.Done()
		barista.PrepareOrder(order, orders)
	}()

	// Wait for PrepareOrder to complete
	go func() {
		wg.Wait()
		close(orders)
	}()

	// Check if the order is received in the channel
	select {
	case receivedOrder := <-orders:
		if receivedOrder != order {
			t.Errorf("Expected order %+v, but got %+v", order, receivedOrder)
		}
	case <-time.After(7 * time.Second): // Adjust timeout as needed
		t.Error("Order was not received in the channel within the expected time")
	}
}
