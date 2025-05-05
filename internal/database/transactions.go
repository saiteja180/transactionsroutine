package database

import (
	"context"
	"transactionroutine/models"
)

func (d *Database) CreateTransaction(ctx context.Context, transaction models.Transaction) (string, error) {
	query := `insert into transaction (account_id, operation_type_id,amount) values ($1,$2,$3) RETURNING transaction_id`
	var transactionId string
	err := d.Conn.QueryRowContext(ctx, query, transaction.AccountId, transaction.OperationTypeID, transaction.Amount).Scan(&transactionId)
	if err != nil {
		return "", err
	}
	return transactionId, nil
}
