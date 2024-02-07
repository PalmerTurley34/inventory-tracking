package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.RoundedBorder(), true)

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
	inventoryList  list.Model
	commandList    list.Model
	appOutputView  viewport.Model
	loginForm      *huh.Form
	initialForm    *huh.Form
	createUserForm *huh.Form
	page           appPage
	focused        appList
	spinner        spinner.Model
	spinnerActive  bool
	spinnerMsg     string
}

func newModel() model {
	m := model{
		inventoryList: list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 30),
		commandList:   list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 10),
		appOutputView: viewport.New(30, 15),
		focused:       invList,
		page:          initialPage,
		initialForm:   NewInitialForm(),
		spinner:       spinner.New(),
		spinnerActive: false,
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
	case userLoggedInMsg:
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.page = mainPage
	case userCreatedMsg:
		m.spinnerActive = false
		m.spinner = spinner.New()
		m.page = initialPage
		m.initialForm = NewInitialForm()
		initCmd := m.initialForm.Init()
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
		return fmt.Sprintf("%v %v", m.spinner.View(), m.spinnerMsg)
	}
	if m.page == mainPage {
		invListView := docStyle.Render(m.inventoryList.View())
		cmdListView := docStyle.Render(m.commandList.View())
		outputView := docStyle.Render(m.appOutputView.View())
		rightSide := lipgloss.JoinVertical(lipgloss.Right, cmdListView, outputView)
		return lipgloss.JoinHorizontal(lipgloss.Top, invListView, rightSide)
	} else if m.page == loginPage {
		return m.loginForm.View()
	} else if m.page == initialPage {
		return m.initialForm.View()
	} else if m.page == createUserPage {
		return m.createUserForm.View()
	}
	return "Page not implemented yet"
}
