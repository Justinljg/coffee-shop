package coffee

import (
	"testing"

	"github.com/justinljg/coffee-shop/coffee"
)

func TestCoffeeTypeToString(t *testing.T) {
	tests := []struct {
		coffeeType coffee.CoffeeType
		expected   string
	}{
		{coffee.Espresso, "Espresso"},
		{coffee.Latte, "Latte"},
		{coffee.Cappuccino, "Cappuccino"},
		{coffee.CoffeeType(100), "Unknown"},
	}

	for _, tt := range tests {
		result := coffee.CoffeeTypeToString(tt.coffeeType)
		if result != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, result)
		}
	}
}
