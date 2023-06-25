-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_balance(
    user_id bigint PRIMARY KEY,
    balance bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_balance;
-- +goose StatementEnd
