package main

import (
	"log"

	"github.com/juliofilizzola/bank/api"
	"github.com/juliofilizzola/bank/configuration/database"
	"github.com/juliofilizzola/bank/configuration/environment"
	db "github.com/juliofilizzola/bank/internal/db/sqlc"
)

func main() {
	config, err := environment.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbs, err := database.ConnectDatabase(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot load database:", err)
	}

	store := db.NewStore(dbs)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
