## Transaction Routine 
This application manages accounts and transactions.
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

Note the account_id from the response.

B.) Retrieve account details:

Use the account_id obtained in the previous step.

Run the "get account details" HTTP request, providing the account_id.

C.) Create a transaction:

Run the "create transaction" HTTP request.

```