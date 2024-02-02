-- name: CheckOutItem :one
INSERT INTO inventory_check_outs (id, created_at, updated_at, user_id, inventory_item_id, checked_out_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CheckInItem :one
UPDATE inventory_check_outs
SET checked_in_at = NOW(),
    updated_at = NOW()
WHERE inventory_item_id = $1 
AND user_id = $2 
AND checked_in_at IS NULL
RETURNING *;