package coffee

import (
	"testing"

	"github.com/justinljg/coffee-shop/coffee"
)

func TestOrderCreation(t *testing.T) {
	customerID := 1
	coffeeType := coffee.Latte

	order := coffee.Order{CustomerID: customerID, CoffeeType: coffeeType}

	if order.CustomerID != customerID {
		t.Errorf("Expected customer ID %d, got %d", customerID, order.CustomerID)
	}

	if order.CoffeeType != coffeeType {
		t.Errorf("Expected coffee type %v, got %v", coffeeType, order.CoffeeType)
	}
}
