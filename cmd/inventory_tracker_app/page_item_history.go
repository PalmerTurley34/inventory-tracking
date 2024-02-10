package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

var historyColumns = []table.Column{
	{Title: "User", Width: 10},
	{Title: "Checked Out", Width: 20},
	{Title: "Checked In", Width: 20},
}

func (m model) updateItemHistoryPage(msg tea.Msg) (model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.Type == tea.KeyEsc {
			m.page = mainPage
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.historyTable, cmd = m.historyTable.Update(msg)
	return m, cmd
}
