package main

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type appList int

const (
	invList appList = iota
	cmdList
)

type appPage int

const (
	loginPage appPage = iota
	initialPage
	mainPage
	createUserPage
	createItemPage
)

type model struct {
	// models
	inventoryList list.Model
	commandList   list.Model
	appOutputView viewport.Model
	// forms
	loginForm      *huh.Form
	initialForm    *huh.Form
	createUserForm *huh.Form
	createItemForm *huh.Form
	// app state
	page      appPage
	focused   appList
	userInfo  db.User
	headerMsg string
	// loading spinner config
	spinner       spinner.Model
	spinnerActive bool
	spinnerMsg    string
	// http clinet
	client *http.Client
	// styles
	headerStyle lipgloss.Style
}

func newModel() model {
	m := model{
		inventoryList: list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 40),
		commandList:   list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 20),
		appOutputView: viewport.New(30, 15),
		focused:       invList,
		page:          initialPage,
		headerMsg:     "Login to begin",
		initialForm:   NewInitialForm(),
		client:        &http.Client{Timeout: 10 * time.Second},
		headerStyle:   successHeaderStyle,
	}
	m.inventoryList.Title = "Toy Box"
	m.commandList.Title = "Commands"
	m.commandList.SetItems(m.DefaultCommands())
	m.resetSpinner()
	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.initialForm.Init(), m.getAllInventoryItems)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case errMsg:
		m.headerMsg = fmt.Sprintf("Error encountered: %v", msg.err)
		m.headerStyle = failureHeaderStyle

	case allInventoryItemsMsg:
		m.inventoryList.SetItems(msg.items)

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

	case userLoggedOutMsg:
		m.userInfo = db.User{}
		m.initialForm = NewInitialForm()
		m.page = initialPage
		m.headerMsg = "Successfully logged out!"
		m.headerStyle = successHeaderStyle
		return m, m.initialForm.Init()

	case itemCreateSuccessMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Successfully created item: %v", msg.item.Name)
		m.headerStyle = successHeaderStyle
		m.resetSpinner()
		m.inventoryList.InsertItem(len(m.inventoryList.Items()), msg.item)

	case itemCreateFailureMsg:
		m.page = mainPage
		m.headerMsg = fmt.Sprintf("Error creating item: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.resetSpinner()

	case startItemCreationMsg:
		m.createItemForm = NewCreateItemForm()
		m.page = createItemPage
		m.headerMsg = "Creating New Item..."
		m.headerStyle = loadingHeaderStyle
		return m, m.createItemForm.Init()
	}

	// Update lists and forms depending on what page you're on //
	if m.page == mainPage {
		return m.updateMainPage(msg)
	}
	if m.page == loginPage {
		return m.updateLoginPage(msg)
	}
	if m.page == initialPage {
		return m.updateInitialPage(msg)
	}
	if m.page == createUserPage {
		return m.updateCreateUserPage(msg)
	}
	if m.page == createItemPage {
		return m.updateCreateItemPage(msg)
	}
	return m, nil
}

func (m model) View() string {
	if m.spinnerActive {
		return loadingHeaderStyle.Render(fmt.Sprintf("%v %v", m.spinner.View(), m.spinnerMsg))
	}
	header := m.headerStyle.Render(fmt.Sprintf("///%v//////", m.headerMsg))
	var pageView string
	if m.page == mainPage {
		pageView = m.getMainPageView()
	} else if m.page == loginPage {
		pageView = m.loginForm.View()
	} else if m.page == initialPage {
		pageView = m.initialForm.View()
	} else if m.page == createUserPage {
		pageView = m.createUserForm.View()
	} else if m.page == createItemPage {
		pageView = m.createItemForm.View()
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, pageView)
}

func (m *model) resetSpinner() {
	m.spinnerActive = false
	m.spinner = spinner.New(spinner.WithSpinner(spinner.Meter))
}
