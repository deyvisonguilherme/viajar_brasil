package models

import "database/sql"
import _ "github.com/lib/pq" // import drivers para postgresql

var db *sql.DB

// InitDB inicializa a conexão com o banco de dados.
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
