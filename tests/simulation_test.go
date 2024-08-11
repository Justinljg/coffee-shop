package cafe

import (
	"context"
	"fmt"
	"github.com/justinljg/coffee-shop/cafe"
	"sync"
	"testing"
)

func TestCoffeeShopSimulation_RaceConditions(t *testing.T) {
	// Increase the number of customers and iterations for a more thorough race condition check
	numCustomers := 1000
	numIterations := 10
	cafe.TestSleepMultiplier = 0.00001

	for i := 0; i < numIterations; i++ {
		// Setup channels and baristas
		customers := make(chan cafe.Customer, numCustomers)
		orders := make(chan cafe.Order, numCustomers)
		baristas := []cafe.Barista{{ID: 1}, {ID: 2}}

		// Use a WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup

		// Simulate customer arrivals
		wg.Add(1)
		go func() {
			defer wg.Done()
			cafe.SimulateCustomerArrivals(context.Background(), customers, numCustomers)
			close(customers)
		}()

		// Baristas prepare orders
		for _, barista := range baristas {
			wg.Add(1)
			go func(b cafe.Barista) {
				defer wg.Done()
				for customer := range customers {
					order := cafe.Order{CustomerID: customer.ID, CoffeeType: cafe.CoffeeType(b.ID % 3)} // Use barista ID for predictable coffee type
					b.PrepareOrder(order, orders)
				}
			}(barista)
		}

		// Collect and count orders
		completedOrders := make(map[int]int) // Map to track orders per customer
		go func() {
			wg.Wait()
			close(orders)
		}()

		for order := range orders {
			completedOrders[order.CustomerID]++
		}

		// Verify that each customer received exactly one order
		for i := 1; i <= numCustomers; i++ {
			if completedOrders[i] != 1 {
				t.Errorf("Iteration %d: Customer %d received %d orders, expected 1", i, i, completedOrders[i])
			}
		}

		fmt.Printf("Iteration %d completed successfully.\n", i+1)
	}
}
