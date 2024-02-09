package main

import (
	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
)

type InventoryItem struct {
	db.InventoryItem
}

func (i InventoryItem) Title() string       { return i.Name }
func (i InventoryItem) Description() string { return i.DueAt.String() }
func (i InventoryItem) FilterValue() string { return i.Name }
