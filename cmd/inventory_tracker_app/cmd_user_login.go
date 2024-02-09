package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	tea "github.com/charmbracelet/bubbletea"
)

type loginFailMsg struct {
	err error
}

type loginSucessMsg struct {
	userInfo db.User
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
