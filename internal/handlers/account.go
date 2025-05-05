package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"transactionroutine/models"
)

func (s *Service) HandleAccountCreation(w http.ResponseWriter, r *http.Request) {
	var req models.Account
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponse(w, nil, http.StatusBadRequest, nil)
		return
	}
	// validate document number
	if req.DocumentNumber == "" {
		WriteResponse(w, nil, http.StatusBadRequest, errors.New("document number is empty"))
		return
	}
	accId, err := s.Db.CreateAccount(ctx, req.DocumentNumber)
	if err != nil {
		WriteResponse(w, nil, http.StatusInternalServerError, err)
		return
	}
	WriteResponse(w, accId, http.StatusCreated, nil)
	return

}

func (s *Service) HandleGetAccountDetails(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	ctx := r.Context()
	accountId := queryParams["accountId"][0]
	if accountId == "" {
		WriteResponse(w, nil, http.StatusBadRequest, errors.New("account id is required"))
		return
	}

	accDetails, err := s.Db.GetAccountDetails(ctx, accountId)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			WriteResponse(w, nil, http.StatusNotFound, nil)
			return
		}
		WriteResponse(w, nil, http.StatusInternalServerError, err)
		return
	}
	WriteResponse(w, accDetails, http.StatusOK, nil)
	return

}
