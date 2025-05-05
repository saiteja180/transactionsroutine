## Transaction Routine 
This application manages accounts and transactions.


## Note
```
1. Added test cases for rest apis which covers db mock and rest apis
2. Added goose migrations folder to automatically run sql scripts

```

## make sure docker is installed in your system

## Folder structure 

## Transactionroutine 
```
Transactionroutine/
├── db/             # Contains SQL scripts for database table creation and initialization.
├── internal/       # Contains the core application logic.
├── models/         # Defines the data structures (schemas) for accounts, transactions, etc.
├── payload/        # Defines the structure of HTTP request payloads.
├── docker-compose/  # Contains the Docker Compose configuration for setting up the PostgreSQL database.
├── go.mod          # Go module definition and dependencies.
├── main.go         # The service starting point (main application entry point).
└── run.sh          # Shell script to start the API and its dependencies.

```


# How to start the service

## Running the API and Testing the Endpoints
```
1. Run the API: Execute the run.sh script:

chmod +x run.sh

./run.sh

2. Test the endpoints: Navigate to the payload folder and execute the HTTP requests as follows:

A.) Create an account:

Run the "create account" HTTP request.
or use curl cmd

curl -X POST http://localhost:8080/api/v1/accounts \
  -H "Content-Type: application/json" \
  -d '{"document_number":"test12345"}'


Note the account_id from the response.

B.) Retrieve account details:

Use the account_id obtained in the previous step.

Run the "get account details" HTTP request, providing the account_id.

or use curl cmd 
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "account_id":"80cb16fc-35e9-4cf6-ad74-e7dc337eb556",
    "operation_type_id": "1",
    "amount": 90
  }'


C.) Create a transaction:

Run the "create transaction" HTTP request.

or use curl cmd
curl "http://localhost:8080/api/v1/accounts?accountId=80cb16fc-35e9-4cf6-ad74-e7dc337eb556"


```
