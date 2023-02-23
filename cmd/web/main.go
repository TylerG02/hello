package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Create a new type
type application struct {
}

func main() {
	//create a flag for speifing the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("GUEST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	//create an instance of the connection pool
	app := &application{}

	defer db.Close()
	log.Println("database connection pool established")
	//Create a customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err) //should not reach here
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a contect to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
