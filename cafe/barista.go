package cafe

import (
	"fmt"
	"time"
)

// Barista with ID to prepare coffee
type Barista struct {
	ID int
}

// PrepareOrder simulates a barista preparing a coffee order.
// It logs the start and end of the preparation, simulates the preparation time,
// and sends the completed order to the provided channel.
func (barista *Barista) PrepareOrder(order Order, orders chan<- Order) {
	// Log the start of the preparation process.
	// Print the barista's ID, the type of coffee being prepared, and the customer ID.
	fmt.Printf("Barista %d starts preparing %s for Customer %d.\n", barista.ID, CoffeeTypeToString(order.CoffeeType), order.CustomerID)

	// Simulate the time taken to prepare the coffee.
	// The preparation time is determined by the type of coffee.
	time.Sleep(CoffeePreparationTimes[order.CoffeeType])

	// Log the completion of the preparation process.
	// Print the barista's ID, the type of coffee, and the customer ID.
	fmt.Printf("Barista %d finishes %s for Customer %d.\n", barista.ID, CoffeeTypeToString(order.CoffeeType), order.CustomerID)

	// Send the completed order to the orders channel.
	// This allows other parts of the program to receive and process the completed order.
	orders <- order
}
