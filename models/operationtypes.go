package models

type OperationsTypes struct {
	OperationTypeID string        `json:"operation_type_id"`
	Description     string        `json:"description"`
	TransactionType OperationType `json:"transaction_type"`
}
type OperationType string

const (
	OperationTypeDebit  OperationType = "debit"
	OperationTypeCredit OperationType = "credit"
)
