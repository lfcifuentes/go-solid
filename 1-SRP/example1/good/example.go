package goodsrp

import "time"

// Order struct represents an order in the system
type Order struct {
	ID          int
	CustomerID  int
	TotalAmount float64
	CreatedAt   time.Time
}

// OrderValidator interface defines the contract for validating orders
// This interface is responsible for validating the order data
type OrderValidator interface {
	Validate(order Order) bool
}

// OrderPrinter interface defines the contract for printing orders
// This interface is responsible for printing the order data
type OrderPrinter interface {
	Print(order Order)
}

// OrderSaver interface defines the contract for saving orders
// This interface is responsible for saving the order data to the database
type OrderSaver interface {
	Save(order Order)
}
