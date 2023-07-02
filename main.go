package main

import (
	"log"

	"github.com/juliofilizzola/bank/api"
	"github.com/juliofilizzola/bank/initializers"
	db "github.com/juliofilizzola/bank/internal/db/sqlc"
)

func init() {
	initializers.Env()
	initializers.ConnectDatabase()
}

func main() {
	store := db.NewStore(initializers.DB)
	server := api.NewServer(store)

	err := server.Start(initializers.PORT)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
