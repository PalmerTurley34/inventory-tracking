package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateLoginPage(msg tea.Msg) (model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.Type == tea.KeyEsc {
			m.page = initialPage
			m.initialForm = NewInitialForm()
			return m, m.initialForm.Init()
		}
	}
	var cmds []tea.Cmd
	form, cmd := m.loginForm.Update(msg)
	cmds = append(cmds, cmd)
	if f, ok := form.(*huh.Form); ok {
		m.loginForm = f
	}
	if m.loginForm.State == huh.StateCompleted {
		m.spinnerActive = true
		m.spinnerMsg = "Logging In..."
		cmds = append(cmds, m.loginUserCmd, m.spinner.Tick)
		m.page = loadingPage
	}
	return m, tea.Batch(cmds...)
}
