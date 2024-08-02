package cafe

import (
	"time"
)

// CoffeeType represents different types of coffee.
type CoffeeType int

// Coffeetype use iota to enumerate coffee type
const (
	Espresso CoffeeType = iota
	Latte
	Cappuccino
)

// CoffeePreparationTimes show the time needed for each type of coffee
var CoffeePreparationTimes = map[CoffeeType]time.Duration{
	Espresso:   2 * time.Second,
	Latte:      3 * time.Second,
	Cappuccino: 4 * time.Second,
}

// CoffeeTypeToString convert Coffee type to strings
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
