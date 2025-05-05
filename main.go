package main

import (
	"context"
	"embed"
	"github.com/gorilla/mux"
	"github.com/pressly/goose/v3"
	"log"
	"net/http"
	"transactionroutine/internal/database"
	"transactionroutine/internal/handlers"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	ctx := context.Background()
	router := mux.NewRouter()
	db, err := database.NewDbConnection()
	if err != nil {
		log.Fatal("error in connecting to DB", err)
	}

	// run sql scripts before starting the api
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("error in setting postgres dialect", err)
	}

	if err := goose.Up(db.Conn, "db/migrations"); err != nil {
		log.Fatal("error in setting up migrations", err)
	}

	svc, err := handlers.NewService(ctx, db)
	if err != nil {
		log.Fatal("error in creating service", err)
	}

	RegisterRoutes(router, svc)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("error in starting http server", err)
	}
}

func RegisterRoutes(r *mux.Router, svc handlers.TransactionsRoutine) {
	r.HandleFunc("/api/v1/transactions", svc.HandlerTransactionCreation).Methods("POST")
	r.HandleFunc("/api/v1/accounts", svc.HandleAccountCreation).Methods("POST")
	r.HandleFunc("/api/v1/accounts", svc.HandleGetAccountDetails).Methods("GET")

}
