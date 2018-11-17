package database_test

import (
	"flag"
	"fmt"
	"net"
	"os"
	"testing"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mguid65/osb-website/server/database"
)

func TestMySQLDB(t *testing.T) {
	t.Parallel()

	user := flag.String("dbuser", "mwalto7", "the database user")
	host := flag.String("dbhost", "127.0.0.1", "the database address")
	port := flag.String("dbport", "3306", "the database port")
	name := flag.String("dbname", "osb_test", "the database name")
	flag.Parse()

	fmt.Fprint(os.Stderr, "DB Password: ")
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		t.Fatalf("could not read password: %v\n", err)
	}
	fmt.Fprintln(os.Stderr)

	db, err := database.Connect(*user, string(pass), net.JoinHostPort(*host, *port), *name)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testUserDB(t, db)
	testResultsDB(t, db)
	testSpecsDB(t, db)
}
