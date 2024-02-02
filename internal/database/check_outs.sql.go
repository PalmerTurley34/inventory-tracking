// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: check_outs.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkInItem = `-- name: CheckInItem :one
UPDATE inventory_check_outs
SET checked_in_at = NOW(),
    updated_at = NOW()
WHERE inventory_item_id = $1 
AND user_id = $2 
AND checked_in_at IS NULL
RETURNING id, user_id, inventory_item_id, created_at, updated_at, checked_out_at, checked_in_at
`

type CheckInItemParams struct {
	InventoryItemID uuid.UUID
	UserID          uuid.UUID
}

func (q *Queries) CheckInItem(ctx context.Context, arg CheckInItemParams) (InventoryCheckOut, error) {
	row := q.db.QueryRowContext(ctx, checkInItem, arg.InventoryItemID, arg.UserID)
	var i InventoryCheckOut
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.InventoryItemID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CheckedOutAt,
		&i.CheckedInAt,
	)
	return i, err
}

const checkOutItem = `-- name: CheckOutItem :one
INSERT INTO inventory_check_outs (id, created_at, updated_at, user_id, inventory_item_id, checked_out_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_id, inventory_item_id, created_at, updated_at, checked_out_at, checked_in_at
`

type CheckOutItemParams struct {
	ID              uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserID          uuid.UUID
	InventoryItemID uuid.UUID
	CheckedOutAt    time.Time
}

func (q *Queries) CheckOutItem(ctx context.Context, arg CheckOutItemParams) (InventoryCheckOut, error) {
	row := q.db.QueryRowContext(ctx, checkOutItem,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.InventoryItemID,
		arg.CheckedOutAt,
	)
	var i InventoryCheckOut
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.InventoryItemID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CheckedOutAt,
		&i.CheckedInAt,
	)
	return i, err
}