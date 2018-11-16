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

	"golang.org/x/crypto/ssh/terminal"

	"github.com/mwalto7/osb-website/server/database"
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

	db, err := database.New(*user, string(passwd), net.JoinHostPort(*host, *port), *name)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	handler := http.NewServeMux()

	handler.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Listing all users!")

		results, err := db.ListResults()
		if err != nil {
			fmt.Fprintln(w, err)
		}
		for _, result := range results {
			fmt.Fprintf(w, "%v\n", result)
		}
	})
//        svr := &http.Server{Addr: ":443", Handler: handler}
//        go func() {
//                fmt.Println("Listening on http:", svr.Addr)
//
//                if err := svr.ListenAndServeTLS("path to key.pem", "path to key.key"); err != nil && err != http.ErrServerClosed {
	svr := &http.Server{Addr: "127.0.0.1:8080", Handler: handler}
	go func() {
		fmt.Println("Listening on http:", svr.Addr)

		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
