package database_test

import (
	"flag"
	"net"
	"testing"

	"github.com/mguid65/osb-website/server/database"
)

var (
	user string
	pass string
	host string
	port string
	name string
)

func init() {
	flag.StringVar(&user, "dbuser", "mwalto7", "the database user")
	flag.StringVar(&pass, "dbpass", "", "database password")
	flag.StringVar(&host, "dbhost", "127.0.0.1", "the database address")
	flag.StringVar(&port, "dbport", "3306", "the database port")
	flag.StringVar(&name, "dbname", "osb_test", "the database name")
}

func TestMySQLDB(t *testing.T) {
	t.Parallel()

	flag.Parse()

	db, err := database.Connect(user, string(pass), net.JoinHostPort(host, port), name)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testUserDB(t, db)
	testResultsDB(t, db)
	// testSpecsDB(t, db)
}
