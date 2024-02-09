package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateCreateUserPage(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.createUserForm.Update(msg)
	cmds = append(cmds, cmd)
	if f, ok := form.(*huh.Form); ok {
		m.createUserForm = f
	}
	if m.createUserForm.State == huh.StateCompleted {
		m.spinnerActive = true
		m.spinnerMsg = "Creating New User..."
		m.page = loadingPage
		cmds = append(cmds, m.createUserCmd, m.spinner.Tick)
	}
	return m, tea.Batch(cmds...)
}
