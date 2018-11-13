package database

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

// User represents the Users MySQL table.
type User struct {
	ID       int64  // user's ID
	Name     string // user's username
	Email    string // user's email
	Password string // user's hashed password
}

// scanUser returns a user from a database row.
func scanUser(s rowScanner) (*User, error) {
	var (
		id       int64
		name     string
		email    string
		password string
	)
	if err := s.Scan(&id, &name, &email, &password); err != nil {
		return nil, err
	}
	user := &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	return user, nil
}
