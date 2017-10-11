package db

import "database/sql"

func NewConnection() sql.DB {
	return sql.DB{}
}

func GetRows(db sql.DB) sql.Rows {
	return sql.Rows{}
}

func Exec(db sql.DB) bool {
	return false
}
