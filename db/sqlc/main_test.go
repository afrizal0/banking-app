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

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Cannot connect db", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
