package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
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

	router := mux.NewRouter()

	router.HandleFunc("/results", handlers.ListResults(db))
	router.HandleFunc("/results/user/{id:[0-9]+}", handlers.ListResultsCreatedBy(db))
	router.HandleFunc("/results/{id:[0-9]+}", handlers.GetResult(db))
	router.HandleFunc("/results/submit", handlers.AddResult(db))
	router.HandleFunc("/results/delete/{id:[0-9]+}", handlers.DeleteResult(db))
	router.HandleFunc("/results/update/{id:[0-9]+}", handlers.UpdateResult(db))

	svr := &http.Server{
		Addr:    "127.0.0.1:443",
		Handler: router,
	}

	var (
		certFile = "/home/osbadmin/cert/key.pem"
		keyFile  = "/home/osbadmin/cert/key.key"
	)
	go func() {
		fmt.Printf("Listening on https://%v\n", svr.Addr)

		if err := svr.ListenAndServeTLS(certFile, keyFile); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	errc := make(chan error, 1)
	go func() {
		defer close(errc)

		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		fmt.Fprintln(os.Stderr, " Shutting down server...")

		if err := svr.Shutdown(ctx); err != nil {
			errc <- fmt.Errorf("could not shut down server within 30s: %v", err)
		}
	}()

	if err := <-errc; err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server successfully shut down.")
}
