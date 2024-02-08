package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type command struct {
	name        string
	description string
	cmd         tea.Cmd
}

func (c command) Title() string       { return c.name }
func (c command) Description() string { return c.description }
func (c command) FilterValue() string { return c.name + c.description }

func (m model) DefaultCommands() []list.Item {
	return []list.Item{
		command{
			name:        "Logout",
			description: "Logout and return to the home page",
			cmd:         m.logoutUserCmd,
		},

		command{
			name:        "Exit",
			description: "Closes the program",
			cmd:         tea.Quit,
		},

		command{
			name:        "Create Item",
			description: "Create a new item and add it to the list",
			cmd:         startItemCreationCmd,
		},
	}
}

func (m model) SelectedItemCommands() []list.Item {
	return []list.Item{
		command{
			name:        "Delete Item",
			description: "Deletes the selected item forever",
		},

		command{
			name:        "History",
			description: "Show history of the selected item",
		},

		command{
			name:        "Check Out Item",
			description: "Checks out the selected item",
		},
	}
}

func (m model) InventoryItemSelectedCommands() []list.Item {
	return []list.Item{
		command{
			name:        "Check In Item",
			description: "Checks in the selected item",
		},
	}
}
