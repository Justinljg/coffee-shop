package cafe

// Order represents a coffee order placed by a customer.
type Order struct {
	CustomerID int        // ID of the customer who placed the order
	CoffeeType CoffeeType // Type of coffee ordered
}
