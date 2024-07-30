package cafe

import (
	"testing"

	"github.com/justinljg/coffee-shop/cafe"
)

func TestCoffeeTypeToString(t *testing.T) {
	tests := []struct {
		coffeeType cafe.CoffeeType
		expected   string
	}{
		{cafe.Espresso, "Espresso"},
		{cafe.Latte, "Latte"},
		{cafe.Cappuccino, "Cappuccino"},
		{cafe.CoffeeType(100), "Unknown"},
	}

	for _, tt := range tests {
		result := cafe.CoffeeTypeToString(tt.coffeeType)
		if result != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, result)
		}
	}
}
