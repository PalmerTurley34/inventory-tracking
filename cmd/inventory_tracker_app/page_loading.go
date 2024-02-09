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
		m.page = loginPage
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
		m.initialForm = NewInitialForm()
		m.page = initialPage
		return m, m.initialForm.Init()

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
		m.resetCommands()
		m.toyBoxList.RemoveItem(m.toyBoxList.Cursor())

	case itemDeleteFailureMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Error deleting item: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()
		m.resetCommands()

	case itemCheckOutSuccessMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Successfully checked out %s", msg.invItem.Name)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.resetCommands()
		m.toyBoxList.SetItem(m.toyBoxList.Cursor(), msg.toyBoxItem)
		m.inventoryList.InsertItem(len(m.inventoryList.Items()), msg.invItem)

	case itemCheckOutFailureMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Error checking out: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()
		m.resetCommands()
	}
	return m, nil
}

func (m *model) resetSpinner() {
	m.spinnerActive = false
	m.spinner = spinner.New(spinner.WithSpinner(spinner.Pulse))
}
