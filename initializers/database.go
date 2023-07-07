package initializers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB
var Err error
var DBs *gorm.DB

// ConnectDatabase todo remove connectDatabase
func ConnectDatabase() {
	DB, Err = sql.Open(DbDriver, UrlDatabase)

	if Err != nil {
		log.Fatal("cannot connect to db:", Err)
	}
}

func Database() {
	dsn := DNS
	var err error
	DBs, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// fmt.Println(db)
}
