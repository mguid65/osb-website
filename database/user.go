package database

// User represents the Users MySQL table.
type User struct {
	ID       int64  // the user ID
	Name     string // user's username
	Email    string // user's email
	Password string // user's hashed password
}

// UserDatabase provides thread-safe access to a database of users.
type UserDatabase interface {
	// ListUsers returns a list of all users.
	ListUsers() ([]*User, error)

	// GetUser retrieves a user by its id.
	GetUser(id int64) (*Result, error)

	// AddUser saves a given user.
	AddUser(res *Result) error

	// DeleteUser deletes a user with the given id.
	DeleteUser(id int64) error

	// UpdateUser updates a given user.
	UpdateUser(res *Result) error
}
