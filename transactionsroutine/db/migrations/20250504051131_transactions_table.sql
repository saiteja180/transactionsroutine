-- +goose Up
-- +goose StatementBegin
CREATE TABLE  if not exists transaction (
                         transaction_id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
                         account_id uuid,
                         operation_type_id varchar(1),
                         amount decimal,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
