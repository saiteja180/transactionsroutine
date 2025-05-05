package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"transactionroutine/internal/cache"
	"transactionroutine/models"
)

func (s *Service) HandlerTransactionCreation(w http.ResponseWriter, r *http.Request) {
	var req models.Transaction
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, nil, http.StatusBadRequest, err)
		return
	}

	// validate account details before creating transaction
	_, err = s.Db.GetAccountDetails(ctx, req.AccountId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			WriteResponse(w, nil, http.StatusNotFound, nil)
			return
		}
		WriteResponse(w, nil, http.StatusInternalServerError, err)
		return
	}

	err = s.validateTransactionDetails(req)
	if err != nil {
		WriteResponse(w, nil, http.StatusBadRequest, err)
		return
	}

	// check the type of transaction
	req.Amount = convertAmount(req.Amount, req.OperationTypeID)

	accId, err := s.Db.CreateTransaction(ctx, req)
	if err != nil {
		WriteResponse(w, nil, http.StatusInternalServerError, err)
		return
	}
	WriteResponse(w, accId, http.StatusCreated, nil)
	return

}

func convertAmount(amt float64, operationType string) float64 {

	t, ok := cache.OperationIdCache[operationType]

	if !ok {
		// if not found then insert as it is
		return amt
	}

	amt = math.Abs(amt)
	if t == models.OperationTypeDebit {
		amt = amt * -1
	}

	return amt

}
func (s *Service) validateTransactionDetails(req models.Transaction) error {
	if req.AccountId == "" {
		return errors.New("account Id is required")
	}
	if req.Amount == 0 {
		return errors.New("amount is required")
	}
	if req.OperationTypeID == "" {
		return errors.New("operation Type Id is required")
	}
	return nil
}
