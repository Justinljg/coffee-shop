package cafe

import (
	"time"
)

// CoffeeType represents the type of coffee being ordered.
// It is defined as an enumeration using iota.
type CoffeeType int

const (
	// Espresso represents a type of coffee.
	Espresso CoffeeType = iota
	// Latte represents a type of coffee.
	Latte
	// Cappuccino represents a type of coffee.
	Cappuccino
)

// CoffeePreparationTimes maps each CoffeeType to its preparation time.
// The duration values simulate the time taken to prepare each type of coffee.
var CoffeePreparationTimes = map[CoffeeType]time.Duration{
	Espresso:   2 * time.Second, // Espresso takes 2 seconds to prepare
	Latte:      3 * time.Second, // Latte takes 3 seconds to prepare
	Cappuccino: 4 * time.Second, // Cappuccino takes 4 seconds to prepare
}

// CoffeeTypeToString converts a CoffeeType to its corresponding string representation.
// This function helps in printing and logging the type of coffee being prepared or ordered.
func CoffeeTypeToString(coffeeType CoffeeType) string {
	switch coffeeType {
	case Espresso:
		return "Espresso"
	case Latte:
		return "Latte"
	case Cappuccino:
		return "Cappuccino"
	default:
		return "Unknown"
	}
}
