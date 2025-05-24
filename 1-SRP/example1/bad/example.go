package badsrp

// Bad example of Single Responsibility Principle
// This example using a order struct to do multiple things
// 1. It is responsible for order data
// 2. It is responsible for order validation
// 3. It is responsible for order printing
// 4. It is responsible for order saving
// 5. It is responsible for order sending

import (
	"fmt"
	"time"
)

// Order struct represents an order in the system
type Order struct {
	ID          int
	CustomerID  int
	TotalAmount float64
	CreatedAt   time.Time
}

// Validate validates the order data
// This method is responsible for validating the order data
// It checks if the customer ID is valid and if the total amount is greater than zero
// If the order is valid, it returns true, otherwise it returns false
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both order data and order validation
func (o *Order) Validate() bool {
	// Validate the order data
	if o.CustomerID == 0 {
		return false
	}
	if o.TotalAmount <= 0 {
		return false
	}
	return true
}

// Print prints the order data
// This method is responsible for printing the order data
// It prints the order ID, customer ID, total amount, and created at date
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both order data and order printing
func (o *Order) Print() {
	fmt.Printf("Order ID: %d\n", o.ID)
	fmt.Printf("Customer ID: %d\n", o.CustomerID)
	fmt.Printf("Total Amount: %.2f\n", o.TotalAmount)
	fmt.Printf("Created At: %s\n", o.CreatedAt.Format(time.RFC3339))
}

// Save saves the order data to the database
// This method is responsible for saving the order data to the database
// It simulates saving the order data to the database
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both order data and order saving
func (o *Order) Save() {
	fmt.Println("Order saved to database")
}

// Send sends the order data to the customer
// This method is responsible for sending the order data to the customer
// It simulates sending the order data to the customer
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both order data and order sending
func (o *Order) Send() {
	fmt.Println("Order sent to customer")
}
