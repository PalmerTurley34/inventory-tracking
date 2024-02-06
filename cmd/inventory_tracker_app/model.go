package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.RoundedBorder(), true)

type appList int

const (
	invList appList = iota
	cmdList
)

type model struct {
	inventoryList list.Model
	commandList   list.Model
	appOutputView viewport.Model
	cmdInput      textinput.Model
	focused       appList
}

func newModel() model {
	m := model{
		inventoryList: list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 30),
		commandList:   list.New([]list.Item{}, list.NewDefaultDelegate(), 30, 10),
		appOutputView: viewport.New(30, 15),
		cmdInput:      textinput.New(),
		focused:       invList,
	}
	m.inventoryList.Title = "Toy Box"
	m.commandList.Title = "Commands"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	}

	var listCmd tea.Cmd
	if m.focused == invList {
		m.inventoryList, listCmd = m.inventoryList.Update(msg)
	} else {
		m.commandList, listCmd = m.commandList.Update(msg)
	}
	return m, listCmd
}

func (m model) View() string {
	invListView := docStyle.Render(m.inventoryList.View())
	cmdListView := docStyle.Render(m.commandList.View())
	outputView := docStyle.Render(m.appOutputView.View())
	rightSide := lipgloss.JoinVertical(lipgloss.Right, cmdListView, outputView)
	return lipgloss.JoinHorizontal(lipgloss.Top, invListView, rightSide)
}
