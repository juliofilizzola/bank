package main

import (
	"log"

	"github.com/juliofilizzola/bank/initializers"
	"github.com/juliofilizzola/bank/internal/db/models"
)

func init() {
	initializers.Env()
	initializers.Database()
}

func main() {
	err := initializers.DBs.AutoMigrate(&models.Account{}, &models.Transfers{})
	if err != nil {
		log.Fatal(err)
	}
}
