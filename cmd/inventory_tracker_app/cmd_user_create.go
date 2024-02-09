package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	tea "github.com/charmbracelet/bubbletea"
)

type userCreateSuccessMsg struct {
	userInfo db.User
}
type userCreateFailMsg struct {
	err error
}

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
