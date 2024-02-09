package main

import (
	"fmt"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
)

type inventoryItem struct {
	db.InventoryItem
}

func (i inventoryItem) Title() string { return i.Name }
func (i inventoryItem) Description() string {
	return fmt.Sprintf("Due At: %s", i.DueAt.Format("02 Jan 06 15:04"))
}
func (i inventoryItem) FilterValue() string { return i.Name }
