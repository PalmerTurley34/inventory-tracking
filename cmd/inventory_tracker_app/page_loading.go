package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateLoadingPage(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case loginSucessMsg:
		m.headerMsg = fmt.Sprintf("Logged in as %s", msg.userInfo.Username)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.userInfo = msg.userInfo
		m.page = mainPage
		return m, m.getAllInventoryItemsCmd

	case loginFailMsg:
		m.headerMsg = fmt.Sprintf("Error logging in: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()
		m.loginForm = NewLoginForm()
		return m, m.loginForm.Init()

	case userCreateSuccessMsg:
		m.headerMsg = fmt.Sprintf("Successfully created user %s", msg.userInfo.Username)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.page = initialPage
		m.initialForm = NewInitialForm()
		return m, m.initialForm.Init()

	case userCreateFailMsg:
		m.headerMsg = fmt.Sprintf("Error creating user: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()
		m.createUserForm = NewCreateUserForm()
		return m, m.createUserForm.Init()

	case itemCreateSuccessMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Successfully created item: %v", msg.item.Name)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.toyBoxList.InsertItem(len(m.toyBoxList.Items()), msg.item)

	case itemCreateFailureMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Error creating item: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()

	case itemDeleteSuccessMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Successfully deleted item: %v", msg.item.Name)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.toyBoxList.RemoveItem(m.toyBoxList.Cursor())
		m.commandList.SetItems(m.DefaultCommands())
		m.isItemSelected = false
		m.focused = toyBoxList

	case itemDeleteFailureMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Error deleting item: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()
		m.commandList.SetItems(m.DefaultCommands())
		m.isItemSelected = false
		m.focused = toyBoxList
	}
	return m, nil
}

func (m *model) resetSpinner() {
	m.spinnerActive = false
	m.spinner = spinner.New(spinner.WithSpinner(spinner.Meter))
}
