package cafe

import (
	"testing"
	"time"

)

// This is a unit test. Unit tests should belong in the same package as the code.
// coffee_test.go should be in the same folder as coffee.go.
func TestCoffeeTypeToString(t *testing.T) {
	tests := []struct {
		coffeeType CoffeeType
		expected   string
	}{
		{Espresso, "Espresso"},
		{Latte, "Latte"},
		{Cappuccino, "Cappuccino"},
		{CoffeeType(999), "Unknown"}, // Testing an undefined CoffeeType
	}

	for _, test := range tests {
		result := CoffeeTypeToString(test.coffeeType)
		if result != test.expected {
			t.Errorf("CoffeeTypeToString(%d) = %s; want %s", test.coffeeType, result, test.expected)
		}
	}
}

// This is a unit test. Unit tests should belong in the same package as the code.
func TestCoffeePreparationTimes(t *testing.T) {
	tests := []struct {
		coffeeType CoffeeType
		expected   time.Duration
	}{
		{Espresso, 2 * time.Second},
		{Latte, 3 * time.Second},
		{Cappuccino, 4 * time.Second},
	}

	for _, test := range tests {
		result, ok := CoffeePreparationTimes[test.coffeeType]
		if !ok {
			t.Errorf("CoffeePreparationTimes does not contain entry for CoffeeType %d", test.coffeeType)
			continue
		}
		if result != test.expected {
			t.Errorf("CoffeePreparationTimes[%d] = %v; want %v", test.coffeeType, result, test.expected)
		}
	}
}
