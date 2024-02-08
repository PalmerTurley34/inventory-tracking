package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (m model) updateConfirmationPage(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.confirmationForm.Update(msg)
	cmds = append(cmds, cmd)
	if f, ok := form.(*huh.Form); ok {
		m.confirmationForm = f
	}
	if m.confirmationForm.State == huh.StateCompleted {
		m.spinnerActive = true
		m.spinnerMsg = "Deleting Item..."
		cmds = append(cmds, m.deleteInventoryItemCmd, m.spinner.Tick)
	}
	return m, tea.Batch(cmds...)
}
