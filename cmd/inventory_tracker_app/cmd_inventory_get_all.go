package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type allInventoryItemsMsg struct {
	items []list.Item
}

func (m model) getAllInventoryItemsCmd() tea.Msg {
	req, _ := http.NewRequest(
		"GET",
		"http://localhost:8080/v1/users/inventory",
		nil,
	)
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", m.userInfo.ApiKey))
	response, err := m.client.Do(req)
	if err != nil {
		return errMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		errResp := struct {
			Error string `json:"error"`
		}{}
		json.NewDecoder(response.Body).Decode(&errResp)
		return errMsg{fmt.Errorf(errResp.Error)}
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
