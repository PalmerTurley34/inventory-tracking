-- name: LogCheckOut :one
INSERT INTO inventory_check_outs (id, created_at, updated_at, user_id, inventory_item_id, checked_out_at)
VALUES ($1, NOW(), NOW(), $2, $3, NOW())
RETURNING *;

-- name: CheckOutItem :one
UPDATE inventory_items
SET checked_out_at = NOW(),
    due_at = NOW() + INTERVAL '24 hours', 
    user_id = $2
WHERE id = $1
RETURNING *;

-- name: LogCheckIn :one
UPDATE inventory_check_outs
SET checked_in_at = NOW(),
    updated_at = NOW()
WHERE inventory_item_id = $1 
AND user_id = $2 
AND checked_in_at IS NULL
RETURNING *;

-- name: CheckInItem :one
UPDATE inventory_items
SET checked_out_at = NULL, user_id = NULL, due_at = NULL, checked_in_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetItemHistory :many
select inventory_items.name, users.username, inventory_check_outs.checked_out_at, inventory_check_outs.checked_in_at
from inventory_check_outs
join inventory_items on inventory_check_outs.inventory_item_id = inventory_items.id
join users on inventory_check_outs.user_id = users.id
where inventory_check_outs.inventory_item_id = $1
order by inventory_check_outs.checked_out_at desc;