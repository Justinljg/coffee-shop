package cafe

import (
	"testing"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

// This is a unit test. Unit tests should belong in the same package as the code.
func TestPrepareOrder(t *testing.T) {
	// Create a barista and an order
	barista := cafe.Barista{ID: 1}
	order := cafe.Order{
		CustomerID: 1,
		CoffeeType: cafe.Espresso, // Use a valid CoffeeType constant
	}

	// Create a channel to receive the order
	// TODO: improve naming of channel.
	// Can be clearer if this is for new orders the barista has received or completed orders.
	orders := make(chan cafe.Order)

	// Run PrepareOrder in a separate goroutine
	go func() {
		barista.PrepareOrder(order, orders)
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
