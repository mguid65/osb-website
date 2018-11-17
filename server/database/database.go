package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

// OSBDatabase provides thread-safe access to users, results, and specs.
type OSBDatabase interface {
	ResultDatabase
	SpecsDatabase
	UserDatabase

	// Close closes the database connection.
	Close() error
}

// Connect establishes a tcp connection to the database.
func Connect(user, passwd, addr, dbName string) (OSBDatabase, error) {
	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = passwd
	cfg.Net = "tcp"
	cfg.Addr = addr
	cfg.DBName = dbName

	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}
	return &mysqlDB{conn: conn, statements: make(map[string]*sql.Stmt)}, nil
}

type mysqlDB struct {
	conn       *sql.DB
	statements map[string]*sql.Stmt
}

// Ensure mysqlDB implements the OSBDatabse interface.
var _ OSBDatabase = &mysqlDB{}

// rowScanner is implemented by sql.Row and sql.Rows.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// newStmt ensures a statement is created, prepared, and stored only once.
func newStmt(db *mysqlDB, once *sync.Once, name, query string) (prepared *sql.Stmt, err error) {
	once.Do(func() {
		if prepared, err = db.conn.Prepare(query); err == nil {
			db.statements[name] = prepared
		}
	})
	if err != nil {
		return nil, fmt.Errorf("db: preapre %s: %v", name, err)
	}
	stmt, ok := db.statements[name]
	if !ok {
		return nil, fmt.Errorf("db: %s not found", name)
	}
	return stmt, nil
}

var prepListResults sync.Once

// ListResults returns a list of all results.
func (db *mysqlDB) ListResults() ([]*Result, error) {
	listResults, err := newStmt(
		db,
		&prepListResults,
		"listResults",
		`SELECT * FROM Results`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := listResults.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*Result
	for rows.Next() {
		result, err := scanResult(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

var listResultsCreatedByOnce sync.Once

// ListResultsCreatedBy returns a list of results created by a user with the given id.
func (db *mysqlDB) ListResultsCreatedBy(id int64) ([]*Result, error) {
	listResultsCreatedBy, err := newStmt(
		db,
		&listResultsCreatedByOnce,
		"listResultsCreatedBy",
		`SELECT * FROM Results WHERE user_id = ?`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := listResultsCreatedBy.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*Result
	for rows.Next() {
		result, err := scanResult(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		results = append(results, result)
	}
	return results, nil
}

var getResultOnce sync.Once

// GetResult retrieves a result by its id.
func (db *mysqlDB) GetResult(id int64) (*Result, error) {
	getResult, err := newStmt(
		db,
		&getResultOnce,
		"getResult",
		`SELECT * FROM Results WHERE user_id = ?`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := scanResult(getResult.QueryRowContext(ctx, id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return result, nil
}

var addResultOnce sync.Once

// AddResult saves a given result.
func (db *mysqlDB) AddResult(res *Result) error {
	addResult, err := newStmt(
		db,
		&addResultOnce,
		"addResult",
		`INSERT INTO Results(user_id, specs_id, results) VALUES(?, ?, ?)`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = addResult.ExecContext(ctx, res.UserID, res.SpecsID, res.Results)
	return err
}

var deleteResultOnce sync.Once

// DeleteResult deletes a result with the given id.
func (db *mysqlDB) DeleteResult(id int64) error {
	deleteResult, err := newStmt(
		db,
		&deleteResultOnce,
		"deleteResult",
		`DELETE FROM Results WHERE result_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = deleteResult.ExecContext(ctx, id)
	return err
}

var updateResultOnce sync.Once

// UpdateResult updates a given result.
func (db *mysqlDB) UpdateResult(res *Result) error {
	updateResult, err := newStmt(
		db,
		&updateResultOnce,
		"updateResult",
		`UPDATE Results SET result = ? WHERE result_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = updateResult.ExecContext(ctx, res.Results, res.ID)
	return err
}

var listSpecsOnce sync.Once

// ListSpecs returns a list of all specs.
func (db *mysqlDB) ListSpecs() ([]*Specs, error) {
	listSpecs, err := newStmt(
		db,
		&listSpecsOnce,
		"listSpecs",
		`SELECT * FROM Specs`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := listSpecs.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specs []*Specs
	for rows.Next() {
		spec, err := scanSpecs(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

var listSpecsCreatedByOnce sync.Once

// ListSpecsCreatedBy returns a list of specs created by a user with the given id.
func (db *mysqlDB) ListSpecsCreatedBy(id int64) ([]*Specs, error) {
	listSpecsCreatedBy, err := newStmt(
		db,
		&listSpecsCreatedByOnce,
		"listSpecsCreatedBy",
		`SELECT * FROM Specs WHERE specs_id = ?`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := listSpecsCreatedBy.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specs []*Specs
	for rows.Next() {
		spec, err := scanSpecs(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

var getSpecsOnce sync.Once

// GetSpecs retrieves specs by its id.
func (db *mysqlDB) GetSpecs(id int64) (*Specs, error) {
	getSpecs, err := newStmt(
		db,
		&getSpecsOnce,
		"getSpecs",
		`SELECT * FROM Specs WHERE specs_id = ?`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	spec, err := scanSpecs(getSpecs.QueryRowContext(ctx, id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return spec, nil

}

var addSpecsOnce sync.Once

// AddSpecs saves the given specs.
func (db *mysqlDB) AddSpecs(specs *Specs) error {
	addSpecs, err := newStmt(
		db,
		&addSpecsOnce,
		"addSpecs",
		`INSERT INTO Specs(result_id, sys_info) VALUES(?, ?)`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = addSpecs.ExecContext(ctx, specs.ResultID, specs.SysInfo)
	return err
}

var deleteSpecsOnce sync.Once

// DeleteSpecs deletes the specs with the given id.
func (db *mysqlDB) DeleteSpecs(id int64) error {
	deleteSpecs, err := newStmt(
		db,
		&deleteSpecsOnce,
		"deleteSpecs",
		`DELETE FROM Specs WHERE specs_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = deleteSpecs.ExecContext(ctx, id)
	return err
}

var updateSpecsOnce sync.Once

// UpdateSpecs updates the given specs.
func (db *mysqlDB) UpdateSpecs(specs *Specs) error {
	updateSpecs, err := newStmt(
		db,
		&updateSpecsOnce,
		"updateSpecs",
		`UPDATE Specs SET sys_info = ? WHERE specs_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = updateSpecs.ExecContext(ctx, specs.SysInfo, specs.ID)
	return err
}

var listUsersOnce sync.Once

// ListUsers returns a list of all users.
func (db *mysqlDB) ListUsers() ([]*User, error) {
	listUsers, err := newStmt(
		db,
		&listUsersOnce,
		"listUsers",
		`SELECT * FROM Users`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := listUsers.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

var getUserOnce sync.Once

// GetUser retrieves a user by its id.
func (db *mysqlDB) GetUser(id int64) (*User, error) {
	getUser, err := newStmt(
		db,
		&getUserOnce,
		"getUser",
		`SELECT * from Users WHERE user_id = ?`,
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := scanUser(getUser.QueryRowContext(ctx, id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return user, nil
}

var getUserByCredentialsOnce sync.Once

// GetUserByCredentials returns a user with the matching username and password.
func (db *mysqlDB) GetUserByCredentials(username, password string) (*User, error) {
	getUserByCredentials, err := newStmt(
		db,
		&getUserByCredentialsOnce,
		"getUserByCredentials",
		``, // TODO: fill in query
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := scanUser(getUserByCredentials.QueryRowContext(ctx, username, password))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return user, nil
}

var addUserOnce sync.Once

// AddUser saves a given user.
func (db *mysqlDB) AddUser(user *User) error {
	addUser, err := newStmt(
		db,
		&addUserOnce,
		"addUser",
		`INSERT INTO Users(username, email, passwd) VALUES(?, ?, ?)`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = addUser.ExecContext(ctx, user.Name, user.Email, user.Password); err != nil {
		return fmt.Errorf("mysql: add user: %v", err)
	}
	return nil
}

var deleteUserOnce sync.Once

// DeleteUser deletes a user with the given id.
func (db *mysqlDB) DeleteUser(id int64) error {
	deleteUser, err := newStmt(
		db,
		&deleteUserOnce,
		"deleteUser",
		`DELETE FROM Users WHERE user_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = deleteUser.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("mysql: delete user: %v", err)
	}
	return nil
}

var updateUserOnce sync.Once

// UpdateUser updates a given user.
func (db *mysqlDB) UpdateUser(user *User) error {
	updateUser, err := newStmt(
		db,
		&updateUserOnce,
		"updateUser",
		`UPDATE Users SET username = ?, email = ?, passwd = ? WHERE user_id = ?`,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err = updateUser.ExecContext(ctx, user.Name, user.Email, user.Password, user.ID); err != nil {
		return fmt.Errorf("mysql: update user: %v", err)
	}
	return nil
}

func (db *mysqlDB) Close() error {
	for _, stmt := range db.statements {
		stmt.Close()
	}
	return db.conn.Close()
}
