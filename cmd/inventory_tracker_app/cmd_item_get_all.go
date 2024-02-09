package main

import (
	"encoding/json"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type allInventoryItemsMsg struct {
	items []list.Item
}

type errMsg struct {
	err error
}

func (m model) getAllInventoryItemsCmd() tea.Msg {
	response, err := m.client.Get("http://localhost:8080/v1/inventory_items")
	if err != nil {
		return errMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errMsg{err}
	}
	items := []inventoryItem{}
	err = json.NewDecoder(response.Body).Decode(&items)
	if err != nil {
		return errMsg{err}
	}
	listItems := []list.Item{}
	for _, i := range items {
		listItems = append(listItems, i)
	}
	return allInventoryItemsMsg{listItems}
}