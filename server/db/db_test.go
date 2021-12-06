package db

import (
	"log"
	"testing"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	connection := &Connection{
		DbName:     "balancersdb",
		Schema:     "public",
		User:       "postgres",
		Password:   "root",
		Host:       "localhost",
		DisableSSL: true,
	}
	log.Println(connection.ConnectionUrl())
	if connection.ConnectionUrl() != "postgres://postgres:root@localhost/balancersdb?search_path=public&sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
