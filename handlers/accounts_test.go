package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactionroutine/internal/database"
	"transactionroutine/models"
)

func TestService_HandleAccountCreation(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		mockDB         func(sqlmock.Sqlmock)
	}{
		{
			name:           "Successful Account Creation",
			requestBody:    `{"document_number": "12345"}`,
			expectedStatus: http.StatusCreated,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("insert into account").
					WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow("1"))

			},
		},
		{
			name:           "Invalid JSON Request",
			requestBody:    `{"document_number": "12345"`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Document Number",
			requestBody:    `{"document_number": ""}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Database Creation Error",
			requestBody:    `{"document_number": "67890"}`,
			expectedStatus: http.StatusInternalServerError,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("insert into account").
					WillReturnError(sql.ErrConnDone)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock database instance
			db, mock, err := sqlmock.New()
			if err != nil {
				// handle error
			}
			defer db.Close()
			if tt.mockDB != nil {
				tt.mockDB(mock)
			}

			s := &Service{
				Db: &database.Database{
					Conn: db,
				},
			}

			req, err := http.NewRequest("POST", "/accounts", bytes.NewBufferString(tt.requestBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			s.HandleAccountCreation(rr, req)

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, rr.Code)

		})
	}
}

func TestService_HandleGetAccountDetails(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      string
		expectedStatus   int
		expectedResponse models.Account
		mockDB           func(sqlmock.Sqlmock)
	}{
		{
			name:        "Successful Get Account  details",
			requestBody: "12345",
			expectedResponse: models.Account{
				DocumentNumber: "12345",
				ID:             "1",
			},
			expectedStatus: http.StatusOK,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

			},
		},
		{
			name:           "Get Account  details not found",
			requestBody:    "12345",
			expectedStatus: http.StatusNotFound,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnError(sql.ErrNoRows)

			},
		},
		{
			name:           "Empty account id",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Database select Error",
			requestBody:    `{"document_number": "67890"}`,
			expectedStatus: http.StatusInternalServerError,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnError(sql.ErrConnDone)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock database instance
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()
			if tt.mockDB != nil {
				tt.mockDB(mock)
			}

			s := &Service{
				Db: &database.Database{
					Conn: db,
				},
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("/accounts?accountId=%s", tt.requestBody), bytes.NewBufferString(tt.requestBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			s.HandleGetAccountDetails(rr, req)

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, rr.Code)
			var r models.Account
			_ = json.Unmarshal(rr.Body.Bytes(), &r)
			assert.Equal(t, tt.expectedResponse, r)
		})
	}
}
