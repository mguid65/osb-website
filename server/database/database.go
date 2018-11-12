package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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

type mysqlDB struct {
	conn *sql.DB
}

var _ OSBDatabase = &mysqlDB{}

func newMySQLDB(user, passwd, addr, dbName string) (OSBDatabase, error) {
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

// rowScanner is implemented by sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

func scanResult(s rowScanner) (*Result, error) {
	var (
		id      int64
		userID  int64
		specsID int64
		results string
	)
	if err := s.Scan(&id, &userID, &specsID, &results); err != nil {
		return nil, err
	}
	result := &Result{
		ID:      id,
		UserID:  userID,
		SpecsID: specsID,
	}
	err := json.NewDecoder(strings.NewReader(results)).Decode(&result.Results)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListResults returns a list of results.
func (db *mysqlDB) ListResults() ([]*Result, error) {
	list, err := db.conn.Prepare(`SELECT * FROM Results`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listResults: %v", err)
	}
	defer list.Close()

	rows, err := list.Query()
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

func (db *mysqlDB) ListResultsCreatedBy(id int64) ([]*Result, error) {
	listCreatedBy, err := db.conn.Prepare(`SELECT * FROM Results WHERE UserID = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listResults: %v", err)
	}
	defer listCreatedBy.Close()

	rows, err := listCreatedBy.Query(id)
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

func (db *mysqlDB) GetResult(id int64) (*Result, error) {
	getResult, err := db.conn.Prepare(`SELECT * FROM Results WHERE ID = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listResults: %v", err)
	}
	defer getResult.Close()

	result, err := scanResult(getResult.QueryRow(id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return result, nil
}

func (db *mysqlDB) AddResult(res *Result) error {
	panic("implement me")
}

func (db *mysqlDB) DeleteResult(id int64) error {
	panic("implement me")
}

func (db *mysqlDB) UpdateResult(res *Result) error {
	panic("implement me")
}

func (db *mysqlDB) ListSpecs() ([]*Specs, error) {
	panic("implement me")
}

func (db *mysqlDB) ListSpecsCreatedBy(id int64) ([]*Specs, error) {
	panic("implement me")
}

func (db *mysqlDB) GetSpecs(id int64) (Specs, error) {
	panic("implement me")
}

func (db *mysqlDB) AddSpecs(specs *Specs) error {
	panic("implement me")
}

func (db *mysqlDB) DeleteSpecs(id int64) error {
	panic("implement me")
}

func (db *mysqlDB) UpdateSpecs(specs *Specs) error {
	panic("implement me")
}

func (db *mysqlDB) ListUsers() ([]*User, error) {
	panic("implement me")
}

func (db *mysqlDB) GetUser(id int64) (*Result, error) {
	panic("implement me")
}

func (db *mysqlDB) AddUser(res *Result) error {
	panic("implement me")
}

func (db *mysqlDB) DeleteUser(id int64) error {
	panic("implement me")
}

func (db *mysqlDB) UpdateUser(res *Result) error {
	panic("implement me")
}

func (db *mysqlDB) Close() error {
	return db.conn.Close()
}
