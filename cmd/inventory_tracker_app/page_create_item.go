package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateCreateItemPage(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.createItemForm.Update(msg)
	cmds = append(cmds, cmd)
	if f, ok := form.(*huh.Form); ok {
		m.createItemForm = f
	}
	if m.createItemForm.State == huh.StateCompleted {
		m.spinnerActive = true
		m.spinnerMsg = "Creating New Item..."
		m.page = loadingPage
		cmds = append(cmds, m.createItemCmd, m.spinner.Tick)
	}
	return m, tea.Batch(cmds...)
}
