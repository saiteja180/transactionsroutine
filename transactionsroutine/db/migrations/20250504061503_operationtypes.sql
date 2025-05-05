-- +goose Up
-- +goose StatementBegin

CREATE TYPE  transaction_type AS ENUM ('credit', 'debit');

CREATE TABLE if not exists operation_types (
                            operation_type_id integer PRIMARY KEY,
                            Description VARCHAR(100),
                            transaction_type transaction_type,
                             created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO operation_types(operation_type_id, Description,transaction_type) values

        (1,'Normal Purchase','debit'),(2,'Purchase with installments','debit'),
        (3,'Withdrawal','debit'),(4,'Credit Voucher','credit')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS transaction_type;
-- +goose StatementEnd
