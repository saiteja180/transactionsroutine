package database

import (
	"context"
	"transactionroutine/models"
)

func (d *Database) CreateAccount(ctx context.Context, docNumber string) (string, error) {
	query := `insert into account (document_number) values ($1) RETURNING account_id`
	var id string
	err := d.Conn.QueryRowContext(ctx, query, docNumber).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (d *Database) GetAccountDetails(ctx context.Context, id string) (models.Account, error) {
	query := `select account_id,document_number from account where account_id = $1`
	var account models.Account
	err := d.Conn.QueryRowContext(ctx, query, id).Scan(&account.ID, &account.DocumentNumber)
	if err != nil {
		return account, err
	}
	return account, nil
}
