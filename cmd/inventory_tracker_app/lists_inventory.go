package main

import (
	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
)

type inventoryItem struct {
	db.InventoryItem
}

func (i inventoryItem) Title() string       { return i.Name }
func (i inventoryItem) Description() string { return i.DueAt.String() }
func (i inventoryItem) FilterValue() string { return i.Name }
