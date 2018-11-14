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

	"golang.org/x/sync/errgroup"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/mwalto7/osb-website/server/database"
)

func main() {
	user := flag.String("dbuser", "osbadmin", "the database user")
	host := flag.String("dbhost", "127.0.0.1", "the database address")
	port := flag.String("dbport", "3306", "the database port")
	name := flag.String("dbname", "osb_db", "the database name")
	flag.Parse()

	passwd, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(fmt.Errorf("could not read password: %v", err))
	}

	db, err := database.NewMySQLDB(*user, string(passwd), net.JoinHostPort(*host, *port), *name)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	svr := &http.Server{}
	g.Go(func() error {
		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt, os.Kill)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-signals:
			cancel()
			return fmt.Errorf("received signal: %v: %v", sig, svr.Shutdown(ctx))
		}
	})
	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
