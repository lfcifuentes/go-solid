package badsrp

import "fmt"

type User struct {
	ID       int
	Name     string
	Age      int
	Password string
}

func (u *User) Validate() bool {
	// Validate user data
	if u.ID <= 0 {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Age <= 0 {
		return false
	}
	if u.Password == "" {
		return false
	}
	return true
}

// SaveToDatabase saves the user to the database
// This method is responsible for saving the user data to the database
// It simulates saving the user data to the database
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both user data and user saving
func (u *User) SaveToDatabase() {
	fmt.Println("User saved to database")
}

// SendWelcomeEmail sends a welcome email to the user
// This method is responsible for sending a welcome email to the user
// It simulates sending a welcome email to the user
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both user data and user email sending
func (u *User) SendWelcomeEmail() {
	fmt.Println("Sending welcome email")
}

// SendPasswordResetEmail sends a password reset email to the user
// This method is responsible for sending a password reset email to the user
// It simulates sending a password reset email to the user
// This method is a violation of the Single Responsibility Principle
// because it is responsible for both user data and user email sending
func (u *User) ValidatePassword() {
	fmt.Println("Validating password")
}
