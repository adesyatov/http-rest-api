package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "sqlserver://testuser:testuser123@localhost?database=testdb_testing&connection+timeout=30&app+name=http-rest-api"
	}
	os.Exit(m.Run())
}
