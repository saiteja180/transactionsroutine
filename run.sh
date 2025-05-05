docker compose up -d


#sleep 2m

go mod tidy

go build -o transactionsroutine

./transactionsroutine