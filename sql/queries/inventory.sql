-- name: CreateInventoryItem :one
INSERT INTO inventory_items (id, created_at, updated_at, name)
VAlUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllInventoryItems :many
SELECT * FROM inventory_items;