package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mguid65/osb-website/server/database"
	"github.com/mguid65/osb-website/server/handlers"
)

func main() {
	user := flag.String("dbuser", "osbadmin", "the database user")
	host := flag.String("dbhost", "127.0.0.1", "the database address")
	port := flag.String("dbport", "3306", "the database port")
	name := flag.String("dbname", "osb_db", "the database name")
	flag.Parse()

	fmt.Fprint(os.Stderr, "DB Password: ")
	passwd, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("could not read password: %v\n", err)
	}
	fmt.Fprintln(os.Stderr)

	db, err := database.Connect(*user, string(passwd), net.JoinHostPort(*host, *port), *name)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var (
	    addr     = ":443"
	    certFile = "/home/osbadmin/cert/key.pem"
	    keyFile  = "/home/osbadmin/cert/key.key"
	    handler  = handlers.Handler(db)
	)

	fmt.Println("Listening on https://localhost:443/")
	log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, handler))
}
