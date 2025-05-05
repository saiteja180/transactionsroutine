package handlers

import (
	"context"
	"net/http"
	"transactionroutine/internal/cache"
	"transactionroutine/internal/database"
)

type Service struct {
	Db *database.Database
}
type TransactionsRoutine interface {
	HandleGetAccountDetails(w http.ResponseWriter, r *http.Request)
	HandlerTransactionCreation(w http.ResponseWriter, r *http.Request)
	HandleAccountCreation(w http.ResponseWriter, r *http.Request)
}

func NewService(ctx context.Context, db *database.Database) (TransactionsRoutine, error) {
	operatorIdCache, err := db.LoadOperationTypes(ctx)
	if err != nil {
		return nil, err
	}
	for _, val := range operatorIdCache {
		cache.OperationIdCache[val.OperationTypeID] = val.TransactionType
	}

	return &Service{
		Db: db,
	}, nil
}
