package database

import (
	"database/sql"
	"fmt"
	"log"
)

var DBConn *sql.DB

const (
	host = "localhost"
	port = 5433
	user = "postgres"
	dbname = "samarth"
)

func SetupDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
    host, port, user, dbname)

	var err error
	DBConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DBConn.Ping()
	if err != nil {
		log.Fatal("Error connecting to DB.")
	}

	log.Println("DB Connection Successful!")
}
