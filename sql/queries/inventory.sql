-- name: CreateInventoryItem :one
INSERT INTO inventory_items (id, created_at, updated_at, name)
VAlUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllInventoryItems :many
SELECT * FROM inventory_items;

-- name: DeleteInventoryItem :exec
DELETE FROM inventory_items where id = $1;

-- name: GetUserInventory :many
SELECT * FROM inventory_items 
WHERE user_id = $1
ORDER BY due_at DESC;