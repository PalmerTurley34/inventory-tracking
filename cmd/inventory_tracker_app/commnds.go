package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
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
	return userCreatedMsg{user}
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
	time.Sleep(3 * time.Second)
	return userLoggedOutMsg{}
}
