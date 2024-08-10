package cafe

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Barista with ID to prepare coffee
type Barista struct {
	ID int
}

// ServeCustomers starts serving customers.
// It listens to the customers channel, prepares their orders,
// and sends completed orders to the orders channel.
func (barista *Barista) ServeCustomers(ctx context.Context, customers <-chan Customer, orders chan<- Order, wg *sync.WaitGroup, rng *rand.Rand) {
	defer wg.Done()
	for {
		select {
		case customer, ok := <-customers:
			if !ok {
				// Customers channel is closed
				return
			}
			order := Order{
				CustomerID: customer.ID,
				CoffeeType: CoffeeType(rng.Intn(3)), // Random coffee type
			}

			fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.ID, CoffeeTypeToString(order.CoffeeType))

			// Prepare the order.
			barista.PrepareOrder(order, orders)

		case <-ctx.Done():
			fmt.Println("Barista received cancellation signal, stopping.")
			return
		}
	}
}

// PrepareOrder simulates a barista preparing a coffee order.
// It logs the start and end of the preparation, simulates the preparation time,
// and sends the completed order to the provided channel.
func (barista *Barista) PrepareOrder(order Order, orders chan<- Order) {
	fmt.Printf("Barista %d starts preparing %s for Customer %d.\n", barista.ID, CoffeeTypeToString(order.CoffeeType), order.CustomerID)

	// Simulate the time taken to prepare the coffee.
	time.Sleep(CoffeePreparationTimes[order.CoffeeType])

	fmt.Printf("Barista %d finishes %s for Customer %d.\n", barista.ID, CoffeeTypeToString(order.CoffeeType), order.CustomerID)

	// Send the completed order to the orders channel.
	orders <- order
}
