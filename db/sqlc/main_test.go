package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	databaseURL = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
