package initializers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open(DbDriver, UrlDatabase)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
}
