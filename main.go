package main

import (
	"database/sql"
	"log"

	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/api"
	db "github.com/florian-nguyen/tech-school_simple-bank/simple-bank/db/sqlc"
	"github.com/florian-nguyen/tech-school_simple-bank/simple-bank/util"
	_ "github.com/lib/pq" // blind import is necessary to talk with database
)

func main() {

	config, err := util.LoadConfig(".") // config environment variables in root folder.

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot establish connection:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
