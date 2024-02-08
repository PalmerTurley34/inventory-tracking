package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) updateMainPage(msg tea.Msg) (model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.Type {
		case tea.KeyTab:
			if m.focused == invList {
				m.focused = cmdList
			} else {
				m.focused = invList
			}
		case tea.KeyEnter:
			if m.focused == cmdList {
				cmd := m.commandList.SelectedItem()
				if command, ok := cmd.(command); ok {
					return m, command.cmd
				}
			}
		}
	}

	var cmd tea.Cmd
	if m.focused == invList {
		m.inventoryList, cmd = m.inventoryList.Update(msg)
	} else {
		m.commandList, cmd = m.commandList.Update(msg)
	}
	return m, cmd
}

func (m model) getMainPageView() string {
	var invListStyle, cmdListStyle lipgloss.Style
	if m.focused == invList {
		invListStyle = focusedListStyle
		cmdListStyle = unFocusedListStyle
	} else {
		invListStyle = unFocusedListStyle
		cmdListStyle = focusedListStyle
	}
	invListView := invListStyle.Render(m.inventoryList.View())
	cmdListView := cmdListStyle.Render(m.commandList.View())
	outputView := unFocusedListStyle.Render(m.appOutputView.View())
	rightSide := lipgloss.JoinVertical(lipgloss.Right, cmdListView, outputView)
	return lipgloss.JoinHorizontal(lipgloss.Top, invListView, rightSide)
}
