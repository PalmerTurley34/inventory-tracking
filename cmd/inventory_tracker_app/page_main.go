package main

import (
	"fmt"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) updateMainPage(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
				return m, nil
			}
			if m.focused == invList && len(m.inventoryList.Items()) > 0 {
				m.commandList.SetItems(m.InventorySelectedCommands())
				m.isItemSelected = true
				m.focused = cmdList
				return m, nil
			}
		case tea.KeyEsc:
			if m.isItemSelected {
				m.isItemSelected = false
				m.commandList.SetItems(m.DefaultCommands())
				return m, nil
			}
		}
	case userLoggedOutMsg:
		m.userInfo = db.User{}
		m.initialForm = NewInitialForm()
		m.page = initialPage
		m.headerMsg = "Successfully logged out!"
		m.headerStyle = successHeaderStyle
		return m, m.initialForm.Init()

	case startItemCreationMsg:
		m.createItemForm = NewCreateItemForm()
		m.page = createItemPage
		m.headerMsg = "Creating New Item..."
		m.headerStyle = loadingHeaderStyle
		return m, m.createItemForm.Init()

	case startItemDeletionMsg:
		m.page = confirmPage
		m.confirmationForm = NewConfimationForm()
		m.headerMsg = fmt.Sprintf("Deleting item: %s", m.inventoryList.Title)
		m.headerStyle = loadingHeaderStyle
		return m, m.confirmationForm.Init()
	}

	// let the focused list deal with the msg
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
