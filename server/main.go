package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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

	srv := &http.Server{
		Addr:         "127.0.0.1:443",
		Handler:      handler(db),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on https://", srv.Addr)
	log.Fatal(srv.ListenAndServeTLS("/home/osbadmin/cert/key.pem", "/home/osbadmin/cert/key.key"))
}

func handler(db database.OSBDatabase) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/results", handlers.ListResults(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/user/{id:[0-9]+}", handlers.ListResultsCreatedBy(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/{id:[0-9]+}", handlers.GetResult(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/submit", handlers.AddResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/delete/{id:[0-9]+}", handlers.DeleteResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/update/{id:[0-9]+}", handlers.UpdateResult(db)).Methods(http.MethodPost)
	return r
}
