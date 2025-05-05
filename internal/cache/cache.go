package cache

import "transactionroutine/models"

var OperationIdCache = make(map[string]models.OperationType)
