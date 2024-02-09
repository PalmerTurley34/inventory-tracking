package main

import db "github.com/PalmerTurley34/inventory-tracking/internal/database"

type toyBoxItem struct {
	db.InventoryItem
}

func (i toyBoxItem) Title() string { return i.Name }
func (i toyBoxItem) Description() string {
	if i.CheckedOutAt == nil {
		return successHeaderStyle.Render("	Available")
	}
	return failureHeaderStyle.Render("	Checked Out")
}
func (i toyBoxItem) FilterValue() string { return i.Name }
