package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "rest-api.db")
	if err != nil {
		panic(err)
	}
	return db
}

func GetRows(query string) (*sql.Rows, error) {
	return NewConnection().Query(query)
}

func Exec(query string) bool {
	_, err := NewConnection().Exec(query)
	if err != nil {
		return false
	}
	return true
}
