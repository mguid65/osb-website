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

// New returns a new databse instance.
func New(user, passwd, addr, dbName string) (OSBDatabase, error) {
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
	return &mysqlDB{conn: conn}, nil
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

	rows, err := listResultsCreatedBy.Query(id)
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

	result, err := scanResult(getResult.QueryRow(id))
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
		`INSERT INTO Results(user_id, specs_id, scores) VALUES(?, ?, ?)`,
	)
	if err != nil {
		return err
	}

	_, err = addResult.Exec(res.UserID, res.SpecsID, res.Results)
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

	_, err = deleteResult.Exec(id)
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

	_, err = updateResult.Exec(res.Results, res.ID)
	return err
}

func (db *mysqlDB) ListSpecs() ([]*Specs, error) {
	listSpecs, err := db.conn.Prepare(`SELECT * FROM Specs`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listSpecs: %v", err)
	}
	defer listSpecs.Close()

	rows, err := listSpecs.Query()
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

func (db *mysqlDB) ListSpecsCreatedBy(id int64) ([]*Specs, error) {
	listSpecs, err := db.conn.Prepare(`SELECT * FROM Specs WHERE specs_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listSpecs: %v", err)
	}
	defer listSpecs.Close()

	rows, err := listSpecs.Query(id)
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

func (db *mysqlDB) GetSpecs(id int64) (*Specs, error) {
	getSpecs, err := db.conn.Prepare(`SELECT * FROM Specs WHERE specs_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare getSpecs: %v", err)
	}
	defer getSpecs.Close()

	spec, err := scanSpecs(getSpecs.QueryRow(id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return spec, nil

}

func (db *mysqlDB) AddSpecs(specs *Specs) error {
	addSpecs, err := db.conn.Prepare(`INSERT INTO Specs(result_id, sys_info) VALUES(?, ?)`)
	if err != nil {
		return fmt.Errorf("mysql: prepare addSpecs: %v", err)
	}

	_, err = addSpecs.Exec(specs.ResultID, specs.SysInfo)
	return err
}

func (db *mysqlDB) DeleteSpecs(id int64) error {
	deleteSpecs, err := db.conn.Prepare(`DELETE FROM Specs WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare deleteSpecs: %v", err)
	}

	_, err = deleteSpecs.Exec(id)
	return err
}

func (db *mysqlDB) UpdateSpecs(specs *Specs) error {
	updateSpecs, err := db.conn.Prepare(`UPDATE Specs SET sys_info = ? WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare updateSpecs: %v", err)
	}

	_, err = updateSpecs.Exec(specs.SysInfo, specs.ID)
	return err
}

func (db *mysqlDB) ListUsers() ([]*User, error) {
	listUsers, err := db.conn.Prepare(`SELECT * FROM Users`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listUsers: %v", err)
	}
	defer listUsers.Close()

	rows, err := listUsers.Query()
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

func (db *mysqlDB) GetUser(id int64) (*User, error) {
	getUser, err := db.conn.Prepare(`SELECT * from Users WHERE user_id = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	defer getUser.Close()

	user, err := scanUser(getUser.QueryRow(id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return user, nil
}

func (db *mysqlDB) AddUser(user *User) error {
	addUser, err := db.conn.Prepare(`INSERT INTO Users(username, email, passwd) VALUES(?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("mysql: prepare addUser: %v", err)
	}

	if _, err = addUser.Exec(user.Name, user.Email, user.Password); err != nil {
		return fmt.Errorf("mysql: add user: %v", err)
	}
	return nil
}

func (db *mysqlDB) DeleteUser(id int64) error {
	deleteUser, err := db.conn.Prepare(`DELETE FROM Users WHERE user_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare deleteUser: %v", err)
	}

	if _, err = deleteUser.Exec(id); err != nil {
		return fmt.Errorf("mysql: delete user: %v", err)
	}
	return nil
}

func (db *mysqlDB) UpdateUser(user *User) error {
	updateUser, err := db.conn.Prepare(`UPDATE Users SET username = ?, email = ?, passwd = ? WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare updateUser: %v", err)
	}

	if _, err = updateUser.Exec(user.Name, user.Email, user.Password, user.ID); err != nil {
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
