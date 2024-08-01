package cafe

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

// Customer represents a customer with a unique ID.
type Customer struct {
	ID int
}

// CollectDrink waits for the barista to signal that the order is ready and then prints a message.
// It also sends the completed order to the provided completed channel.
func (customer *Customer) CollectDrink(order Order, orders <-chan Order, completed chan<- Order, wg *sync.WaitGroup) {
	defer wg.Done()

	// Wait for the order to be received from the orders channel
	receivedOrder := <-orders
	fmt.Printf("Collecting Payment for Customer %d\n", receivedOrder.CustomerID)
	// Print a message indicating that the customer has received their order
	fmt.Printf("Customer %d receives their %s and leaves.\n", receivedOrder.CustomerID, CoffeeTypeToString(receivedOrder.CoffeeType))

	// Simulate time taken for the customer to leave
	time.Sleep(1 * time.Second)

	// Send the completed order to the completed channel
	completed <- receivedOrder
}


// SimulateCustomerArrivals continuously generates customer arrivals and sends them to the customers channel.
// This function runs until the parent goroutine signals completion by closing the channel.
// It increments the customer ID for each new customer.
func SimulateCustomerArrivals(customers chan<- Customer, numCustomers int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= numCustomers; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		customers <- Customer{ID: i}
	}
}