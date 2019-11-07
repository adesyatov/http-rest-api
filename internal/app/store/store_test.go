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
		databaseURL = "sqlserver://user:password@localhost?database=dev_db&connection+timeout=30&app+name=http-rest-api"
	}
	os.Exit(m.Run())
}
