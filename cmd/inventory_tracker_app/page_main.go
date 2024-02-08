package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) updateMainPage(msg tea.Msg) (model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.Type {
		case tea.KeyTab:
			if m.isItemSelected {
				// don't let users switch to other lists
				// when an item is selected. Choose a
				// command or esc to get out
				break
			}
			if m.focused == toyBoxList {
				m.focused = cmdList
			} else if m.focused == cmdList {
				m.focused = invList
			} else {
				m.focused = toyBoxList
			}
		case tea.KeyEnter:
			if m.focused == cmdList {
				cmd := m.commandList.SelectedItem()
				if command, ok := cmd.(command); ok {
					return m, command.cmd
				}
			}
			if m.focused == toyBoxList {
				m.commandList.SetItems(m.ToyBoxItemCommands())
				m.isItemSelected = true
				m.focused = cmdList
			}
			if m.focused == invList && len(m.inventoryList.Items()) > 0 {
				m.commandList.SetItems(m.InventorySelectedCommands())
				m.isItemSelected = true
				m.focused = cmdList
			}
		case tea.KeyEsc:
			if m.isItemSelected {
				m.isItemSelected = false
				m.commandList.SetItems(m.DefaultCommands())
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	if m.focused == toyBoxList {
		m.toyBoxList, cmd = m.toyBoxList.Update(msg)
	} else if m.focused == cmdList {
		m.commandList, cmd = m.commandList.Update(msg)
	} else {
		m.inventoryList, cmd = m.inventoryList.Update(msg)
	}
	return m, cmd
}

func (m model) getMainPageView() string {
	var toyBoxListStyle, cmdListStyle, invListStyle lipgloss.Style
	if m.focused == toyBoxList {
		toyBoxListStyle = focusedListStyle
		cmdListStyle = unFocusedListStyle
		invListStyle = unFocusedListStyle
	} else if m.focused == cmdList {
		toyBoxListStyle = unFocusedListStyle
		cmdListStyle = focusedListStyle
		invListStyle = unFocusedListStyle
	} else {
		toyBoxListStyle = unFocusedListStyle
		cmdListStyle = unFocusedListStyle
		invListStyle = focusedListStyle
	}
	toyBoxView := toyBoxListStyle.Render(m.toyBoxList.View())
	cmdListView := cmdListStyle.Render(m.commandList.View())
	inventoryView := invListStyle.Render(m.inventoryList.View())
	rightSide := lipgloss.JoinVertical(lipgloss.Right, cmdListView, inventoryView)
	return lipgloss.JoinHorizontal(lipgloss.Top, toyBoxView, rightSide)
}
