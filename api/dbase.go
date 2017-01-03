package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() {
	db, err := gorm.Open("postgres", "host=localhost user=deyvison dbname=learngo sslmode=disable password=88072762")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}
