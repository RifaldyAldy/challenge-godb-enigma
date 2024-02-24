package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123"
	db_name  = "challenge_godb"
)

func Dbcon() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

	db, _ = sql.Open("postgres", psqlInfo)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return
}
