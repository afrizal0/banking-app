package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:Heilingspeir1@localhost:5432/bankingapp?sslmode=disable"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot connect db", err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
