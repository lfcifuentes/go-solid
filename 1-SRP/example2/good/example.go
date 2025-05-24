package goodsrp

// User struct represents a user in the system
type User struct {
	ID       int
	Name     string
	Age      int
	Password string
}

// UserValidator interface defines the methods for validating user data
type UserValidator interface {
	ValidateUser(user User) bool
	ValidatePassword(user User) bool
}

// UserRepository interface defines the methods for saving and retrieving user data
type UserRepository interface {
	Save(user User)
	FindByID(id int) (User, error)
}

// UserNotifier interface defines the methods for sending notifications to users
type UserNotifier interface {
	SendWelcomeEmail(user User)
	SendPasswordResetEmail(user User)
}
