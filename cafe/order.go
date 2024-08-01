package cafe

// Order represents an order placed by a customer.
type Order struct {
	CustomerID int
	CoffeeType CoffeeType
	Complete   chan Order
}
