package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juliofilizzola/bank/initializers"
)

func init() {
	initializers.Env()
}

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(initializers.DbDriver, initializers.UrlDatabaseTest)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())

}
