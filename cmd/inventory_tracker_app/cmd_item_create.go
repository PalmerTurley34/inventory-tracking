package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type startItemCreationMsg struct{}

type itemCreateSuccessMsg struct {
	item toyBoxItem
}

type itemCreateFailureMsg struct {
	err error
}

func startItemCreationCmd() tea.Msg {
	return startItemCreationMsg{}
}

func (m model) createItemCmd() tea.Msg {
	time.Sleep(time.Second)
	itemName := m.createItemForm.GetString("name")
	jsonStr := fmt.Sprintf(`{"name": "%v"}`, itemName)

	response, err := m.client.Post(
		"http://localhost:8080/v1/inventory_items",
		"application-json",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return itemCreateFailureMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 201 {
		return itemCreateFailureMsg{fmt.Errorf(response.Status)}
	}
	item := toyBoxItem{}
	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return itemCreateFailureMsg{err}
	}
	return itemCreateSuccessMsg{item}
}
