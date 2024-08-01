package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/justinljg/coffee-shop/cafe"
)

func main() {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    customers := make(chan cafe.Customer, 10)
    orders := make(chan cafe.Order, 10)
    collected := make(chan cafe.Order, 10)
    baristas := []cafe.Barista{{ID: 1}, {ID: 2}}
    var wg sync.WaitGroup

    numCustomers := 10

    wg.Add(1)
    go func() {
        defer wg.Done()
        cafe.SimulateCustomerArrivals(customers, numCustomers, &wg)
    }()

    for _, barista := range baristas {
        wg.Add(1)
        go func(b cafe.Barista) {
            defer wg.Done()
            for customer := range customers {
                order := cafe.Order{CustomerID: customer.ID, CoffeeType: cafe.CoffeeType(rng.Intn(3))}
                fmt.Printf("Customer %d arrives and places an order for a %s.\n", customer.ID, cafe.CoffeeTypeToString(order.CoffeeType))
                wg.Add(1)
                go b.PrepareOrder(order, orders, &wg)
            }
        }(barista)
    }

    wg.Add(1)
    go func() {
        defer wg.Done()
        for order := range orders {
            wg.Add(1)
            customer := cafe.Customer{ID: order.CustomerID}
            go customer.CollectDrink(order, orders, collected, &wg)
        }
        close(collected)
    }()

    go func() {
        for completedOrder := range collected {
            fmt.Printf("Completed order for Customer %d\n", completedOrder.CustomerID)
        }
    }()

    wg.Wait()
    close(customers)
    close(orders)

    fmt.Println("Coffee shop closed.")
}
