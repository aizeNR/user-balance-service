-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions(
    id uuid PRIMARY KEY,
    user_id bigint not null,
    amount bigint not null,
    comment text default '',
    operation_date timestamptz not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
