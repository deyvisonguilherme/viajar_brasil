package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}
