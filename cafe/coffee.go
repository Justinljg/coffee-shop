package cafe

import (
	"time"
)

// CoffeeType represents different types of coffee.
type CoffeeType int

const (
	Espresso CoffeeType = iota
	Latte
	Cappuccino
)

var CoffeePreparationTimes = map[CoffeeType]time.Duration{
	Espresso:   2 * time.Second,
	Latte:      3 * time.Second,
	Cappuccino: 4 * time.Second,
}

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


