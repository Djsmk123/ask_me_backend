package main

import (
	"database/sql"
	"log"

	"github.com/djsmk123/askmeapi/api"
	db "github.com/djsmk123/askmeapi/db/sqlc"
	"github.com/djsmk123/askmeapi/utils"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	//run migation
	//runDbMigrations(config.MIGRATIONURL, config.DBSource)
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

func runDbMigrations(migrationUrl string, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)
	if (err) != nil {
		log.Fatal("cannot create migration", err)
		return
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to migrate up", err)
		return
	}
	log.Println("db migration succeeded")
}
