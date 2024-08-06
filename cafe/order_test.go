package cafe

import (
    "testing"
)

func TestOrderInitialization(t *testing.T) {
    // Initialize an order
    order := Order{
        CustomerID: 1,
        CoffeeType: Espresso,
    }

    // Check if the order is correctly initialized
    if order.CustomerID != 1 {
        t.Errorf("Expected CustomerID to be 1, but got %d", order.CustomerID)
    }

    if order.CoffeeType != Espresso {
        t.Errorf("Expected CoffeeType to be Espresso, but got %v", order.CoffeeType)
    }
}

func TestOrderEquality(t *testing.T) {
    // Initialize two orders
    order1 := Order{
        CustomerID: 1,
        CoffeeType: Latte, 
    }

    order2 := Order{
        CustomerID: 1,
        CoffeeType: Latte,
    }

    if order1 != order2 {
        t.Errorf("Expected orders to be equal, but they are not.")
    }
}
