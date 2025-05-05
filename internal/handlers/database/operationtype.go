package database

import (
	"context"
	"transactionroutine/models"
)

func (d *Database) LoadOperationTypes(ctx context.Context) ([]models.OperationsTypes, error) {
	query := `select operation_type_id,transaction_type from operation_types`
	rows, err := d.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var operationTypes []models.OperationsTypes
	for rows.Next() {
		var t models.OperationsTypes
		err = rows.Scan(&t.OperationTypeID, &t.TransactionType)
		if err != nil {
			return nil, err
		}
		operationTypes = append(operationTypes, t)
	}

	return operationTypes, nil
}
