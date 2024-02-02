-- name: CreateInventoryItem :one
INSERT INTO inventory_items (id, created_at, updated_at, name, description)
VAlUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllInventoryItems :many
SELECT * FROM inventory_items;