package database

import (
	"database/sql"
	"fmt"

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

type MysqlDB struct {
	conn *sql.DB
}

var _ OSBDatabase = &MysqlDB{}

func NewMySQLDB(user, passwd, addr, dbName string) (OSBDatabase, error) {
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
	return &MysqlDB{conn: conn}, nil
}

// rowScanner is implemented by sql.Row and sql.Rows.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

func (db *MysqlDB) ListResults() ([]*Result, error) {
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

func (db *MysqlDB) ListResultsCreatedBy(id int64) ([]*Result, error) {
	listCreatedBy, err := db.conn.Prepare(`SELECT * FROM Results WHERE UserID = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare listResultsCreatedBy: %v", err)
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

func (db *MysqlDB) GetResult(id int64) (*Result, error) {
	getResult, err := db.conn.Prepare(`SELECT * FROM Results WHERE ID = ?`)
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare getResult: %v", err)
	}
	defer getResult.Close()

	result, err := scanResult(getResult.QueryRow(id))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not read row: %v", err)
	}
	return result, nil
}

func (db *MysqlDB) AddResult(res *Result) error {
	panic("implement me")
}

func (db *MysqlDB) DeleteResult(id int64) error {
	panic("implement me")
}

func (db *MysqlDB) UpdateResult(res *Result) error {
	panic("implement me")
}

func (db *MysqlDB) ListSpecs() ([]*Specs, error) {
	panic("implement me")
}

func (db *MysqlDB) ListSpecsCreatedBy(id int64) ([]*Specs, error) {
	panic("implement me")
}

func (db *MysqlDB) GetSpecs(id int64) (Specs, error) {
	panic("implement me")
}

func (db *MysqlDB) AddSpecs(specs *Specs) error {
	panic("implement me")
}

func (db *MysqlDB) DeleteSpecs(id int64) error {
	panic("implement me")
}

func (db *MysqlDB) UpdateSpecs(specs *Specs) error {
	panic("implement me")
}

func (db *MysqlDB) ListUsers() ([]*User, error) {
	panic("implement me")
}

func (db *MysqlDB) GetUser(id int64) (*User, error) {
	panic("implement me")
}

func (db *MysqlDB) AddUser(res *User) error {
	panic("implement me")
}

func (db *MysqlDB) DeleteUser(id int64) error {
	panic("implement me")
}

func (db *MysqlDB) UpdateUser(res *User) error {
	panic("implement me")
}

func (db *MysqlDB) Close() error {
	return db.conn.Close()
}
