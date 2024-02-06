-- +goose Up

CREATE TABLE inventory_items (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    checked_out_at TIMESTAMP,
    checked_in_at TIMESTAMP,
    due_at TIMESTAMP,
    user_id UUID REFERENCES users(id)
);

-- +goose Down

DROP TABLE inventory_items;