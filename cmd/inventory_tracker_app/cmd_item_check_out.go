package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	tea "github.com/charmbracelet/bubbletea"
)

type startItemCheckOutMsg struct{}

type itemCheckOutSuccessMsg struct {
	invItem    inventoryItem
	toyBoxItem toyBoxItem
}

type itemCheckOutFailureMsg struct {
	err error
}

func startItemCheckOutCmd() tea.Msg {
	return startItemCheckOutMsg{}
}

func (m model) checkOutItemCmd() tea.Msg {
	time.Sleep(time.Second)
	toCheckOut := m.toyBoxList.SelectedItem()
	item, _ := toCheckOut.(toyBoxItem)
	if item.UserID != nil {
		return itemCheckOutFailureMsg{fmt.Errorf("item is already checked out")}
	}
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://localhost:8080/v1/checkout/%s", item.ID),
		nil,
	)
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", m.userInfo.ApiKey))
	response, err := m.client.Do(req)
	if err != nil {
		return itemCheckOutFailureMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return itemCheckOutFailureMsg{fmt.Errorf(response.Status)}
	}
	toyBoxItem := toyBoxItem{}
	err = json.NewDecoder(response.Body).Decode(&toyBoxItem)
	if err != nil {
		return itemCheckOutFailureMsg{err}
	}

	invItem := inventoryItem{
		db.InventoryItem{
			ID:           toyBoxItem.ID,
			CreatedAt:    toyBoxItem.CreatedAt,
			UpdatedAt:    toyBoxItem.UpdatedAt,
			Name:         toyBoxItem.Name,
			CheckedOutAt: toyBoxItem.CheckedOutAt,
			CheckedInAt:  toyBoxItem.CheckedInAt,
			DueAt:        toyBoxItem.DueAt,
			UserID:       toyBoxItem.UserID,
		},
	}

	if err != nil {
		return itemCheckOutFailureMsg{err}
	}
	return itemCheckOutSuccessMsg{
		invItem:    invItem,
		toyBoxItem: toyBoxItem,
	}
}
