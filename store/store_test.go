package store

// This file also sets up the test fixture.

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var tStore *Store

func init() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(fmt.Sprintf("Failed to create in-memory SQLite3 DB: %v", err))
	}
	tStore, err = NewStoreDB(db)
	if err != nil {
		panic(fmt.Sprintf("Failed to create Store instance: %v", err))
	}
}

func TestNewStore(t *testing.T) {
	_, err := NewStore()
	if err != nil {
		t.Errorf("NewStore() -> (*, %v), want (*, <nil>)", err)
	}
}
