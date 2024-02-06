-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOL NOT NULL DEFAULT FALSE,
    api_key VARCHAR(64) UNIQUE NOT NULL 
    DEFAULT (encode(sha256(random()::text::bytea), 'hex'))
);

-- +goose Down

DROP TABLE users;