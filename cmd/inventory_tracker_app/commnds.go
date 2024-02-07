package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) createUserCmd() tea.Msg {
	time.Sleep(3 * time.Second)
	return userCreatedMsg{}
}

func (m model) loginUserCmd() tea.Msg {
	time.Sleep(3 * time.Second)
	return userLoggedInMsg{}
}

func (m model) logoutUserCmd() tea.Msg {
	time.Sleep(3 * time.Second)
	return userLoggedOutMsg{}
}
