package main

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type startItemDeletionMsg struct{}

type itemDeleteSuccessMsg struct {
	item toyBoxItem
}

type itemDeleteFailureMsg struct {
	err error
}

func startItemDeletionCmd() tea.Msg {
	return startItemDeletionMsg{}
}

func (m model) deleteInventoryItemCmd() tea.Msg {
	time.Sleep(time.Second)
	toDelete := m.toyBoxList.SelectedItem()
	item, _ := toDelete.(toyBoxItem)
	if item.UserID != nil {
		return itemDeleteFailureMsg{fmt.Errorf("cannot delete checked out out item")}
	}
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("http://localhost:8080/v1/inventory_items/%s", item.ID),
		nil,
	)
	if err != nil {
		return itemDeleteFailureMsg{fmt.Errorf("error deleting item: %v", err)}
	}
	response, err := m.client.Do(req)
	if err != nil {
		return itemDeleteFailureMsg{fmt.Errorf("error deleting item: %v", err)}
	}
	if response.StatusCode != 200 {
		return itemDeleteFailureMsg{fmt.Errorf(response.Status)}
	}
	return itemDeleteSuccessMsg{item}
}
