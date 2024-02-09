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
		if !m.confirmationForm.GetBool("confirm") {
			m.page = mainPage
			m.isItemSelected = false
			m.commandList.SetItems(m.DefaultCommands())
			m.focused = toyBoxList
			return m, nil
		}
		m.spinnerActive = true
		m.page = loadingPage
		var cmd tea.Cmd
		switch m.confirmMsg.(type) {
		case startItemDeletionMsg:
			m.spinnerMsg = "Deleteing item..."
			cmd = m.deleteInventoryItemCmd

		case startItemCheckOutMsg:
			m.spinnerMsg = "Checking out item..."
			cmd = m.checkOutItemCmd
		}
		cmds = append(cmds, cmd, m.spinner.Tick)
	}
	return m, tea.Batch(cmds...)
}
