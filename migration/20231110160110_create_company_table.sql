-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS companies;

-- +goose StatementEnd