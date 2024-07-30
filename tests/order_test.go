package cafe

import (
	"testing"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestOrderCreation(t *testing.T) {
	customerID := 1
	coffeeType := cafe.Latte

	order := cafe.Order{CustomerID: customerID, CoffeeType: coffeeType}

	if order.CustomerID != customerID {
		t.Errorf("Expected customer ID %d, got %d", customerID, order.CustomerID)
	}

	if order.CoffeeType != coffeeType {
		t.Errorf("Expected coffee type %v, got %v", coffeeType, order.CoffeeType)
	}
}
