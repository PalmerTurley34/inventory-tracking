package main

import tea "github.com/charmbracelet/bubbletea"

type userLoggedOutMsg struct{}

func (m model) logoutUserCmd() tea.Msg {
	return userLoggedOutMsg{}
}
