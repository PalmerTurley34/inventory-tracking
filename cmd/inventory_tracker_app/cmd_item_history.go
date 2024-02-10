package main

import (
	"encoding/json"
	"fmt"

	db "github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type startItemHistoryMsg struct{}

type itemHistorySuccessMsg struct {
	historyItems []table.Row
}

type itemHistoryFailureMsg struct {
	err error
}

func startItemHistoryCmd() tea.Msg {
	return startItemHistoryMsg{}
}

func (m model) itemHistoryCmd() tea.Msg {
	toFetch := m.toyBoxList.SelectedItem()
	item, _ := toFetch.(toyBoxItem)

	response, err := m.client.Get(
		fmt.Sprintf("http://localhost:8080/v1/inventory_items/history/%v", item.ID),
	)
	if err != nil {
		return itemHistoryFailureMsg{err}
	}

	defer response.Body.Close()
	if response.StatusCode != 200 {
		return itemHistoryFailureMsg{fmt.Errorf(response.Status)}
	}

	history := []db.GetItemHistoryRow{}
	err = json.NewDecoder(response.Body).Decode(&history)
	if err != nil {
		return itemHistoryFailureMsg{err}
	}

	dateFmt := "Jan 02, 2006 15:04"
	rows := []table.Row{}
	for _, row := range history {
		var checkIn string
		if row.CheckedInAt != nil {
			checkIn = row.CheckedInAt.Local().Format(dateFmt)
		}
		rows = append(rows, []string{
			row.Username,
			row.CheckedOutAt.Local().Format(dateFmt),
			checkIn,
		})
	}
	return itemHistorySuccessMsg{rows}
}
