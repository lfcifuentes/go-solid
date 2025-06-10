# Open/Closed Principle (OCP)

## Definition
The Open/Closed Principle states that software entities (classes, modules, functions, etc.) should be open for extension but closed for modification. In Go, this means we should be able to add new functionality without changing existing code.

## Key Concepts in Go
- Use interfaces to define behavior
- Implement new functionality through new types
- Avoid modifying existing, tested code
- Use composition to extend functionality
- Design for extensibility from the start

## Examples in this Directory

### Example 1: Payment Processing
Demonstrates how to handle different payment methods using interfaces:
```go
// Bad approach - needs modification for each new payment type
type PaymentProcessor struct {
    // ...fields
}

func (p *PaymentProcessor) ProcessPayment(method string) error {
    switch method {
    case "credit":
        // process credit card
    case "debit":
        // process debit card
    // Need to modify this function for each new payment method
    }
    return nil
}

// Good approach - open for extension
type PaymentMethod interface {
    Process() error
}

type CreditCardPayment struct{}
type DebitCardPayment struct{}
type CryptoPayment struct{} // New payment method without changing existing code

func (c *CreditCardPayment) Process() error {
    // Process credit card payment
    return nil
}