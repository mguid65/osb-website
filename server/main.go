package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

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

	fmt.Println("Listening on https://localhost:443/")
	log.Fatal(http.ListenAndServeTLS(":443", "/home/osbadmin/cert/key.pem", "/home/osbadmin/cert/key.key", handler(db)))
}

func handler(db database.OSBDatabase) *mux.Router {
	r := mux.NewRouter()
	addUserHandlers(r, db)
	addResultHandlers(r, db)
	addSpecsHandlers(r, db)
	return r
}
func addUserHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/users", handlers.ListUsers(db)).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser(db)).Methods(http.MethodGet)
	r.HandleFunc("/users/register", handlers.AddUser(db)).Methods(http.MethodPost)
	r.HandleFunc("/users/delete/{id:[0-9]+}", handlers.DeleteUser(db)).Methods(http.MethodPost)
	r.HandleFunc("/users/update/{id:[0-9]+}", handlers.UpdateUser(db)).Methods(http.MethodPost)
}

func addResultHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/results", handlers.ListResults(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/user/{id:[0-9]+}", handlers.ListResultsCreatedBy(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/{id:[0-9]+}", handlers.GetResult(db)).Methods(http.MethodGet)
	r.HandleFunc("/results/submit", handlers.AddResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/delete/{id:[0-9]+}", handlers.DeleteResult(db)).Methods(http.MethodPost)
	r.HandleFunc("/results/update/{id:[0-9]+}", handlers.UpdateResult(db)).Methods(http.MethodPost)
}

func addSpecsHandlers(r *mux.Router, db database.OSBDatabase) {
	r.HandleFunc("/specs", handlers.ListSpecs(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/result/{id:[0-9]+}", handlers.ListSpecsCreatedBy(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/{id:[0-9]+}", handlers.GetSpecs(db)).Methods(http.MethodGet)
	r.HandleFunc("/specs/add", handlers.AddSpecs(db)).Methods(http.MethodPost)
	r.HandleFunc("/specs/delete/{id:[0-9]+}", handlers.DeleteSpecs(db)).Methods(http.MethodPost)
	r.HandleFunc("/specs/update/{id:[0-9]+}", handlers.UpdateSpecs(db)).Methods(http.MethodPost)
}
