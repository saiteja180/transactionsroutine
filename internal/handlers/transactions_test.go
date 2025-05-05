package handlers_test

import (
	"bytes"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactionroutine/internal/cache"
	"transactionroutine/internal/database"
	"transactionroutine/internal/handlers"
	"transactionroutine/models"
)

func TestService_HandlerTransactionCreation(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		mockDB         func(sqlmock.Sqlmock)
		mockCache      map[string]models.OperationType
	}{
		{
			name: "Successful Credit Transaction",
			requestBody: `{
				"account_id": "acc123",
				"amount": 100.50,
				"operation_type_id": "1"
			}`,
			expectedStatus: http.StatusCreated,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{"1": models.OperationTypeCredit},
		},
		{
			name: "Successful Debit Transaction",
			requestBody: `{
				"account_id": "acc456",
				"amount": 50.20,
				"operation_type_id": "2"
			}`,
			expectedStatus: http.StatusCreated,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{"2": models.OperationTypeDebit},
		},
		{
			name: "Invalid JSON Request",
			requestBody: `{
				"account_id": "acc789",
				"amount": 25.00,
				"operation_type_id": "3"`, // Missing closing brace
			expectedStatus: http.StatusBadRequest,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{},
		},
		{
			name: "Missing Account ID",
			requestBody: `{
				"amount": 75.00,
				"operation_type_id": "4"
			}`,
			expectedStatus: http.StatusBadRequest,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{},
		},
		{
			name: "Missing Amount",
			requestBody: `{
				"account_id": "acc101",
				"operation_type_id": "5"
			}`,
			expectedStatus: http.StatusBadRequest,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{},
		},
		{
			name: "Missing Operation Type ID",
			requestBody: `{
				"account_id": "acc112",
				"amount": 120.00
			}`,
			expectedStatus: http.StatusBadRequest,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{},
		},
		{
			name: "Database Creation Error",
			requestBody: `{
				"account_id": "acc131",
				"amount": 10.99,
				"operation_type_id": "1"
			}`,
			expectedStatus: http.StatusInternalServerError,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnError(errors.New("database error"))

			},
			mockCache: map[string]models.OperationType{"1": models.OperationTypeCredit},
		},
		{
			name: "Operation Type Not In Cache (Amount Unchanged)",
			requestBody: `{
				"account_id": "acc141",
				"amount": 99.99,
				"operation_type_id": "unknown"
			}`,
			expectedStatus: http.StatusCreated,
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("select account_id,document_number").
					WillReturnRows(sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow("1", "12345"))

				mock.ExpectQuery("insert into transaction").
					WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow("1"))

			},
			mockCache: map[string]models.OperationType{"1": models.OperationTypeCredit, "2": models.OperationTypeDebit},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err)
			}
			defer db.Close()
			if tt.mockDB != nil {
				tt.mockDB(mock)
			}

			cache.OperationIdCache = tt.mockCache

			s := &handlers.Service{
				Db: &database.Database{
					Conn: db,
				},
			}

			// Create a new HTTP request with the test body
			req, err := http.NewRequest("POST", "/transactions", bytes.NewBufferString(tt.requestBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			s.HandlerTransactionCreation(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

		})
	}
}
