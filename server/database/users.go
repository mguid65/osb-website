package database

// UserDatabase provides thread-safe access to a database of users.
type UserDatabase interface {
	// ListUsers returns a list of all users.
	ListUsers() ([]*UserExternal, error)

	// GetUser retrieves a user by its id.
	GetUser(id int64) (*UserExternal, error)

	// GetUserByCredentials returns a user with the matching username and password.
	GetUserByCredentials(user, pass string) (*User, error)

	// AddUser saves a given user.
	AddUser(user *User) (int64, error)

	// DeleteUser deletes a user with the given id.
	DeleteUser(id int64) error

	// UpdateUser updates a given user.
	UpdateUser(user *User) error
}

// User represents the Users MySQL table.
type User struct {
	ID       int64  // user's ID
	Name     string // user's username
	Email    string // user's email
	Password string // user's hashed password
}

// UserExternal represents a public view of a user in system
type UserExternal struct {
        ID       int64  // user's ID
        Name     string // user's username
}

// scanUser returns a user from a database row.
func scanUserExternal(s rowScanner) (*UserExternal, error) {
        var (
                id       int64
                name     string
        )
        if err := s.Scan(&id, &name); err != nil {
                return nil, err
        }
        userExt := &UserExternal{
                ID:       id,
                Name:     name,
        }
        return userExt, nil
}

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
