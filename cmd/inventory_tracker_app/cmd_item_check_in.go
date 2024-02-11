package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type startItemCheckInMsg struct{}

type itemCheckInSuccessMsg struct {
	item toyBoxItem
}

type itemCheckInFailureMsg struct {
	err error
}

func startItemCheckInCmd() tea.Msg {
	return startItemCheckInMsg{}
}

func (m model) itemCheckInCmd() tea.Msg {
	toCheckIn := m.inventoryList.SelectedItem()
	item, _ := toCheckIn.(inventoryItem)
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("http://localhost:8080/v1/inventory_items/checkin/%s", item.ID),
		nil,
	)
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", m.userInfo.ApiKey))
	response, err := m.client.Do(req)
	if err != nil {
		return itemCheckInFailureMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return itemCheckInFailureMsg{fmt.Errorf(response.Status)}
	}
	toyItem := toyBoxItem{}
	err = json.NewDecoder(response.Body).Decode(&toyItem)
	if err != nil {
		return itemCheckInFailureMsg{err}
	}
	return itemCheckInSuccessMsg{toyItem}
}
