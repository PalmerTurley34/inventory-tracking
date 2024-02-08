package main

import (
	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/charmbracelet/bubbles/list"
)

type errMsg struct {
	err error
}

type allInventoryItemsMsg struct {
	items []list.Item
}

type loginFailMsg struct {
	err error
}

type loginSucessMsg struct {
	userInfo db.User
}

type userLoggedOutMsg struct{}

type userCreateFailMsg struct {
	err error
}

type startItemCreationMsg struct{}

type userCreateSuccessMsg struct {
	userInfo db.User
}

type itemCreateSuccessMsg struct {
	item inventoryItem
}

type itemCreateFailureMsg struct {
	err error
}
