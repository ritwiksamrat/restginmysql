package util

import (
	"database/sql"
)

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@/shopping")
	if err != nil {
		panic(err.Error())
	}

	return db
}
