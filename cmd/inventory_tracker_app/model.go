package main

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type appList int

const (
	toyBoxList appList = iota
	cmdList
	invList
)

type appPage int

const (
	initialPage appPage = iota
	loginPage
	mainPage
	createUserPage
	createItemPage
	confirmPage
	loadingPage
)

type updateMethod func(tea.Msg) (model, tea.Cmd)
type viewMethod func() string

type model struct {
	// models
	toyBoxList    list.Model
	commandList   list.Model
	inventoryList list.Model
	// forms
	loginForm        *huh.Form
	initialForm      *huh.Form
	createUserForm   *huh.Form
	createItemForm   *huh.Form
	confirmationForm *huh.Form
	// app state
	page           appPage
	focused        appList
	userInfo       db.User
	headerMsg      string
	headerStyle    lipgloss.Style
	isItemSelected bool
	confirmMsg     tea.Msg
	// loading spinner config
	spinner       spinner.Model
	spinnerActive bool
	spinnerMsg    string
	// http clinet
	client *http.Client
}

func newModel() model {
	m := model{
		toyBoxList:       list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 40),
		commandList:      list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 18),
		inventoryList:    list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 18),
		focused:          toyBoxList,
		page:             initialPage,
		headerMsg:        "Login to begin",
		initialForm:      NewInitialForm(),
		loginForm:        NewLoginForm(),
		createUserForm:   NewCreateUserForm(),
		createItemForm:   NewCreateItemForm(),
		confirmationForm: NewConfimationForm(),
		client:           &http.Client{Timeout: 10 * time.Second},
		headerStyle:      successHeaderStyle,
	}
	m.toyBoxList.Title = "Toy Box"
	m.commandList.Title = "Commands"
	m.inventoryList.Title = "Inventory"
	m.commandList.SetItems(m.DefaultCommands())
	m.resetSpinner()

	return m
}

func (m model) UpdateMethods() map[appPage]updateMethod {
	return map[appPage]updateMethod{
		initialPage:    m.updateInitialPage,
		loginPage:      m.updateLoginPage,
		createUserPage: m.updateCreateUserPage,
		createItemPage: m.updateCreateItemPage,
		confirmPage:    m.updateConfirmationPage,
		loadingPage:    m.updateLoadingPage,
		mainPage:       m.updateMainPage,
	}
}

func (m model) ViewMethods() map[appPage]viewMethod {
	return map[appPage]viewMethod{
		initialPage:    m.initialForm.View,
		loginPage:      m.loginForm.View,
		createUserPage: m.createUserForm.View,
		createItemPage: m.createItemForm.View,
		confirmPage:    m.confirmationForm.View,
		loadingPage: func() string {
			return loadingHeaderStyle.Render(fmt.Sprintf("%v %v", m.spinner.View(), m.spinnerMsg))
		},
		mainPage: m.getMainPageView,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.initialForm.Init(), m.getAllToyBoxItemsCmd)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case errMsg:
		m.headerMsg = fmt.Sprintf("Error encountered: %v", msg.err.Error())
		m.headerStyle = failureHeaderStyle

	case allToyBoxItemsMsg:
		m.toyBoxList.SetItems(msg.items)

	case allInventoryItemsMsg:
		m.inventoryList.SetItems(msg.items)
	}

	// Update lists and forms depending on what page you're on //
	return m.UpdateMethods()[m.page](msg)
}

func (m model) View() string {
	if m.spinnerActive {
		return loadingHeaderStyle.Render(fmt.Sprintf("%v %v", m.spinner.View(), m.spinnerMsg))
	}
	header := m.headerStyle.Render(fmt.Sprintf("///%v//////\n", m.headerMsg))
	pageView := m.ViewMethods()[m.page]()
	return lipgloss.JoinVertical(lipgloss.Left, header, pageView)
}

func (m *model) resetCommands() {
	m.isItemSelected = false
	m.commandList.SetItems(m.DefaultCommands())
	m.commandList.ResetSelected()
	m.focused = toyBoxList
}
