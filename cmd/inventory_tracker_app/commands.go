package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) createUserCmd() tea.Msg {
	time.Sleep(time.Second)
	username := m.createUserForm.GetString("username")
	password := m.createUserForm.GetString("password")
	name := m.createUserForm.GetString("name")
	jsonStr := fmt.Sprintf(
		`{"name": "%v", "username": "%v", "password": "%v"}`,
		name,
		username,
		password,
	)

	response, err := m.client.Post(
		"http://localhost:8080/v1/users",
		"application/json",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return userCreateFailMsg{err}
	}
	defer response.Body.Close()
	if response.StatusCode != 201 {
		return userCreateFailMsg{fmt.Errorf("%v: %v", response.Status, response.Body)}
	}
	user := db.User{}
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return userCreateFailMsg{err}
	}
	return userCreateSuccessMsg{user}
}

func (m model) loginUserCmd() tea.Msg {
	time.Sleep(time.Second)
	username := m.loginForm.GetString("username")
	password := m.loginForm.GetString("password")
	jsonStr := fmt.Sprintf(`{"username": "%v", "password": "%v"}`, username, password)

	response, err := m.client.Post(
		"http://localhost:8080/v1/login",
		"application/json",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return loginFailMsg{err: err}
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return loginFailMsg{fmt.Errorf("login attempt unsuccessful")}
	}
	user := db.User{}
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return loginFailMsg{err}
	}
	return loginSucessMsg{user}
}

func (m model) logoutUserCmd() tea.Msg {
	return userLoggedOutMsg{}
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
	item := inventoryItem{}
	err = json.NewDecoder(response.Body).Decode(&item)
	if err != nil {
		return itemCreateFailureMsg{err}
	}
	return itemCreateSuccessMsg{item}
}

func startItemCreationCmd() tea.Msg {
	return startItemCreationMsg{}
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

func (m model) deleteInventoryItemCmd() tea.Msg {
	time.Sleep(time.Second)
	toDelete := m.inventoryList.SelectedItem()
	item, _ := toDelete.(inventoryItem)
	if item.UserID != nil {
		return itemDeleteFailureMsg{fmt.Errorf("cannot delete checkout out item")}
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

func startItemDeletionCmd() tea.Msg {
	return startItemDeletionMsg{}
}
