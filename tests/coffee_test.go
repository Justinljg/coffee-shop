package cafe

import (
	"testing"
	"time"
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
		{cafe.CoffeeType(999), "Unknown"}, // Testing an undefined CoffeeType
	}

	for _, test := range tests {
		result := cafe.CoffeeTypeToString(test.coffeeType)
		if result != test.expected {
			t.Errorf("CoffeeTypeToString(%d) = %s; want %s", test.coffeeType, result, test.expected)
		}
	}
}

func TestCoffeePreparationTimes(t *testing.T) {
	tests := []struct {
		coffeeType cafe.CoffeeType
		expected   time.Duration
	}{
		{cafe.Espresso, 2 * time.Second},
		{cafe.Latte, 3 * time.Second},
		{cafe.Cappuccino, 4 * time.Second},
	}

	for _, test := range tests {
		result, ok := cafe.CoffeePreparationTimes[test.coffeeType]
		if !ok {
			t.Errorf("CoffeePreparationTimes does not contain entry for CoffeeType %d", test.coffeeType)
			continue
		}
		if result != test.expected {
			t.Errorf("CoffeePreparationTimes[%d] = %v; want %v", test.coffeeType, result, test.expected)
		}
	}
}
