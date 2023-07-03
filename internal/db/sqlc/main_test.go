package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbDriver  = "mysql"
	dbSources = "root:123456@tcp/bank-test?parseTime=true"
)

func TestMain(m *testing.M) {
	var err error
	// todo: add database from test
	testDb, err = sql.Open(dbDriver, dbSources)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())

}
