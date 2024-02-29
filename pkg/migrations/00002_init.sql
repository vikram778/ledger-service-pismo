-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts
(
    account_id SERIAL PRIMARY KEY NOT NULL,
    document_number VARCHAR(20)   NOT NULL,
    created_at timestamptz default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS operation_types
(
    operation_type_id Int PRIMARY KEY NOT NULL,
    description VARCHAR(30)   NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id SERIAL PRIMARY KEY NOT NULL,
    account_id INT  NOT NULL,
    operation_type_id INT NOT NULL,
    amount NUMERIC(10,2),
    event_date timestamptz default CURRENT_TIMESTAMP
);

INSERT INTO operation_types (operation_type_id,description)
    VALUES (1,'Normal Purchase'),
           (2,'Purchase with installments'),
           (3,'Withdrawal'),
           (4,'Credit Voucher');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS operation_types;
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
