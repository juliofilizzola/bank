package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver  = "mysql"
	dbSources = "root:123456@tcp/bank-dev?parseTime=true"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSources)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}
