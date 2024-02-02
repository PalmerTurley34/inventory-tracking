package models

import (
	"time"

	"github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/google/uuid"
)

type InventoryItem struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func DBInventoryItemToResponse(item database.InventoryItem) InventoryItem {
	var description *string
	if item.Description.Valid {
		description = &item.Description.String
	} else {
		description = nil
	}
	return InventoryItem{
		ID:          item.ID,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Name:        item.Name,
		Description: description,
	}
}

func DBInventoryItemsToResponse(items []database.InventoryItem) []InventoryItem {
	retItems := []InventoryItem{}
	for _, item := range items {
		retItems = append(retItems, DBInventoryItemToResponse(item))
	}
	return retItems
}
