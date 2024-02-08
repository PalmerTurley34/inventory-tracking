package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateInitialPage(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.initialForm.Update(msg)
	cmds = append(cmds, cmd)
	if f, ok := form.(*huh.Form); ok {
		m.initialForm = f
	}
	if m.initialForm.State == huh.StateCompleted {
		if m.initialForm.GetString("option") == "Login" {
			m.loginForm = NewLoginForm()
			initCmd := m.loginForm.Init()
			cmds = append(cmds, initCmd)
			m.page = loginPage
		} else {
			m.createUserForm = NewCreateUserForm()
			initCmd := m.createUserForm.Init()
			cmds = append(cmds, initCmd)
			m.page = createUserPage
		}
	}
	return m, tea.Batch(cmds...)
}
