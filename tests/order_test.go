package cafe

import (
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestOrderChannel(t *testing.T) {
	// Create an order with a buffered channel
	orderComplete := make(chan cafe.Order, 1)
	order := cafe.Order{
		CustomerID: 1,
		CoffeeType: cafe.Espresso,
		Complete:   orderComplete,
	}

	// Simulate sending the order to the Complete channel
	go func() {
		order.Complete <- order
	}()

	// Wait for the order to be received from the Complete channel
	select {
	case receivedOrder := <-order.Complete:
		if receivedOrder != order {
			t.Errorf("Expected order %+v, but got %+v", order, receivedOrder)
		}
	case <-time.After(2 * time.Second): // Adjust timeout as needed
		t.Error("Order was not received in the Complete channel within the expected time")
	}
}
