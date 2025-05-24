# Single Responsibility Principle (SRP)

## Definition
The Single Responsibility Principle states that a package, type, or function should have only one reason to change. In Go, this means each component should have a single, well-defined purpose.

## Key Concepts in Go
- Each type/struct should focus on doing one specific thing
- A package should have a single, focused purpose
- Functions should be small and handle a single operation
- Avoid mixing unrelated functionalities in the same type

## Examples in this Directory

### Example 1: Order Processing
Demonstrates how to separate order processing responsibilities using Go interfaces and types:
```go
// Instead of one large type:
type OrderManager struct {
    // Multiple responsibilities
}

// Split into focused types:
type OrderValidator interface {
    Validate(order Order) error
}

type OrderRepository interface {
    Save(order Order) error
}

type OrderNotifier interface {
    Notify(order Order) error
}
```

### Example 2: User Management
Shows separation of concerns using Go packages and interfaces:
- `user/storage`: Handles user data persistence
- `user/auth`: Manages authentication
- `user/notification`: Handles user notifications

## Bad vs Good Implementation in Go

### Bad Implementation
```go
// Single type handling multiple responsibilities
type User struct {
    // ...fields
}

func (u *User) SaveToDatabase() error {}
func (u *User) SendWelcomeEmail() error {}
func (u *User) ValidatePassword() error {}
```

### Good Implementation
```go
type User struct {
    // ...fields
}

type UserStorage interface {
    Save(user User) error
}

type UserNotifier interface {
    SendWelcome(user User) error
}

type UserAuthenticator interface {
    ValidatePassword(user User) error
}
```

## Benefits
1. More idiomatic Go code
2. Better testability using interfaces
3. Easier to implement dependency injection
4. Clear package boundaries
5. Simplified error handling
6. Better alignment with Go's composition over inheritance approach

## How to Run Examples
Each example contains its own README with specific instructions on how to run tests and examples using `go test` and `go run`.