-- +goose Up

CREATE TABLE inventory_check_outs (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    inventory_item_id UUID NOT NULL REFERENCES inventory_items(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    checked_out_at TIMESTAMP NOT NULL,
    checked_in_at TIMESTAMP,
    UNIQUE(user_id, inventory_item_id, checked_in_at)
);

-- +goose Down

DROP TABLE inventory_check_outs;