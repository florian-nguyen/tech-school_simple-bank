package main

import (
	"database/sql"
	"log"

	"github.com/florian-nguyen/golang-training/tech-school/simple-bank/api"
	db "github.com/florian-nguyen/golang-training/tech-school/simple-bank/db/sqlc"

	_ "github.com/lib/pq" // blind import is necessary to talk with database
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:rootroot@localhost:5432/simple-bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot establish connection:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("Cannot start server:", err)
	}
}
