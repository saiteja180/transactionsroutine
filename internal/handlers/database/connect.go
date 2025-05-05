package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "mypassword"
	dbname   = "transactions"
)

type Database struct {
	Conn *sql.DB
}

func NewDbConnection() (*Database, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	return &Database{Conn: db}, nil
}
func connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to Postgre db")

	return db, nil
}
