package main

import (
	"database/sql"
	"log"

	"github.com/djsmk123/askmeapi/api"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfiguration(".")
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	store := db.NewDBSQLExec(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
