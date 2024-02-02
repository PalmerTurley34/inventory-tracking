-- +goose Up

CREATE TABLE inventory_items (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    description TEXT
);

-- +goose Down

DROP TABLE inventory_items;