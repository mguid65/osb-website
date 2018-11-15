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
	addResult, err := db.conn.Prepare(`INSERT INTO Results(user_id, specs_id, scores) VALUES(?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("mysql: prepare addResult: %v", err)
	}

	_, err = addResult.Exec(res.UserID, res.SpecsID, res.Results)
	return err
}

func (db *MysqlDB) DeleteResult(id int64) error {
	deleteResult, err := db.conn.Prepare(`DELETE FROM Results WHERE result_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare deleteResult: %v", err)
	}

	_, err = deleteResult.Exec(id)
	return err
}

func (db *MysqlDB) UpdateResult(res *Result) error {
	updateResult, err := db.conn.Prepare(`UPDATE Results SET result = ? WHERE result_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare updateResult: %v", err)
	}

	_, err = updateResult.Exec(res.Results, res.ID)
	return err
}

func (db *MysqlDB) ListSpecs() ([]*Specs, error) {
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

func (db *MysqlDB) ListSpecsCreatedBy(id int64) ([]*Specs, error) {
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

func (db *MysqlDB) GetSpecs(id int64) (*Specs, error) {
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

func (db *MysqlDB) AddSpecs(specs *Specs) error {
	addSpecs, err := db.conn.Prepare(`INSERT INTO Specs(result_id, sys_info) VALUES(?, ?)`)
	if err != nil {
		return fmt.Errorf("mysql: prepare addSpecs: %v", err)
	}

	_, err = addSpecs.Exec(specs.ResultID, specs.SysInfo)
	return err
}

func (db *MysqlDB) DeleteSpecs(id int64) error {
	deleteSpecs, err := db.conn.Prepare(`DELETE FROM Specs WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare deleteSpecs: %v", err)
	}

	_, err = deleteSpecs.Exec(id)
	return err
}

func (db *MysqlDB) UpdateSpecs(specs *Specs) error {
	updateSpecs, err := db.conn.Prepare(`UPDATE Specs SET sys_info = ? WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare updateSpecs: %v", err)
	}

	_, err = updateSpecs.Exec(specs.SysInfo, specs.ID)
	return err
}

func (db *MysqlDB) ListUsers() ([]*User, error) {
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

func (db *MysqlDB) GetUser(id int64) (*User, error) {
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

func (db *MysqlDB) AddUser(user *User) error {
	addUser, err := db.conn.Prepare(`INSERT INTO Users(username, email, passwd) VALUES(?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("mysql: prepare addUser: %v", err)
	}

	_, err = addUser.Exec(user.Name, user.Email, user.Password)
	return err
}

func (db *MysqlDB) DeleteUser(id int64) error {
	deleteUser, err := db.conn.Prepare(`DELETE FROM Users WHERE user_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare deleteUser: %v", err)
	}

	_, err = deleteUser.Exec(id)
	return err
}

func (db *MysqlDB) UpdateUser(user *User) error {
	updateUser, err := db.conn.Prepare(`UPDATE Users SET username = ?, email = ?, passwd = ? WHERE specs_id = ?`)
	if err != nil {
		return fmt.Errorf("mysql: prepare updateUser: %v", err)
	}

	_, err = updateUser.Exec(user.Name, user.Email, user.Password, user.ID)
	return err
}

func (db *MysqlDB) Close() error {
	return db.conn.Close()
}
