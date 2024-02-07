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
		inventoryList: list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 30),
		commandList:   list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 10),
		appOutputView: viewport.New(30, 15),
		focused:       invList,
		page:          initialPage,
		headerMsg:     "Login to begin",
		initialForm:   NewInitialForm(),
		spinner:       spinner.New(),
		spinnerActive: false,
		client:        &http.Client{Timeout: 10 * time.Second},
		headerStyle:   successHeaderStyle,
	}
	m.inventoryList.Title = "Toy Box"
	m.commandList.Title = "Commands"
	return m
}

func (m model) Init() tea.Cmd {
	return m.initialForm.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyTab:
			if m.focused == invList {
				m.focused = cmdList
			} else {
				m.focused = invList
			}
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case loginSucessMsg:
		m.headerMsg = fmt.Sprintf("Logged in as %s", msg.userInfo.Username)
		m.headerStyle = successHeaderStyle
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.userInfo = msg.userInfo
		m.page = mainPage
	case loginFailMsg:
		m.headerMsg = fmt.Sprintf("Error logging in: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.loginForm = NewLoginForm()
		initCmd := m.loginForm.Init()
		return m, initCmd
	case userCreatedMsg:
		m.headerMsg = fmt.Sprintf("Successfully created user %s", msg.userInfo.Username)
		m.headerStyle = successHeaderStyle
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.page = initialPage
		m.initialForm = NewInitialForm()
		initCmd := m.initialForm.Init()
		return m, initCmd
	case userCreateFailMsg:
		m.headerMsg = fmt.Sprintf("Error creating user: %v", msg.err)
		m.headerStyle = failureHeaderStyle
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.createUserForm = NewCreateUserForm()
		initCmd := m.createUserForm.Init()
		return m, initCmd
	}

	// Update lists and forms depending on what page you're on //
	if m.page == mainPage {
		if m.focused == invList {
			m.inventoryList, cmd = m.inventoryList.Update(msg)
		} else {
			m.commandList, cmd = m.commandList.Update(msg)
		}
	} else if m.page == loginPage {
		var form tea.Model
		form, cmd = m.loginForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.loginForm = f
		}
		if m.loginForm.State == huh.StateCompleted {
			m.spinnerActive = true
			m.spinnerMsg = "Logging In..."
			cmds = append(cmds, m.loginUserCmd)
			cmds = append(cmds, m.spinner.Tick)
		}
	} else if m.page == initialPage {
		var form tea.Model
		form, cmd = m.initialForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.initialForm = f
		}
		if m.initialForm.State == huh.StateCompleted {
			if m.initialForm.GetString("option") == "Login" {
				m.loginForm = NewLoginForm()
				initCmd := m.loginForm.Init()
				cmds = append(cmds, initCmd)
				m.page = loginPage
				return m, tea.Batch(cmds...)
			} else {
				m.createUserForm = NewCreateUserForm()
				initCmd := m.createUserForm.Init()
				cmds = append(cmds, initCmd)
				m.page = createUserPage
				return m, tea.Batch(cmds...)
			}
		}
	} else if m.page == createUserPage {
		var form tea.Model
		form, cmd = m.createUserForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.createUserForm = f
		}
		if m.createUserForm.State == huh.StateCompleted {
			m.spinnerActive = true
			m.spinnerMsg = "Creating New User..."
			cmds = append(cmds, m.createUserCmd)
			cmds = append(cmds, m.spinner.Tick)
		}
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.spinnerActive {
		return loadingHeaderStyle.Render(fmt.Sprintf("%v %v", m.spinner.View(), m.spinnerMsg))
	}
	header := m.headerStyle.Render(fmt.Sprintf("///%v//////", m.headerMsg))
	var pageView string
	if m.page == mainPage {
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
		pageView = lipgloss.JoinHorizontal(lipgloss.Top, invListView, rightSide)
	} else if m.page == loginPage {
		pageView = m.loginForm.View()
	} else if m.page == initialPage {
		pageView = m.initialForm.View()
	} else if m.page == createUserPage {
		pageView = m.createUserForm.View()
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, pageView)
}
