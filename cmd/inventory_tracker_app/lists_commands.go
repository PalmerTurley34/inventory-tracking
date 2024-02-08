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
			name:        "Create Item",
			description: "Add new item to toy box",
			cmd:         startItemCreationCmd,
		},

		command{
			name:        "Logout",
			description: "Return to the home page",
			cmd:         m.logoutUserCmd,
		},

		command{
			name:        "Exit",
			description: "Close the program",
			cmd:         tea.Quit,
		},
	}
}

func (m model) ToyBoxItemCommands() []list.Item {
	return []list.Item{
		command{
			name:        "Delete Item",
			description: "Delete item forever",
			cmd:         startItemDeletionCmd,
		},

		command{
			name:        "History",
			description: "Show item history",
		},

		command{
			name:        "Check Out",
			description: "Add item to inventory",
		},
	}
}

func (m model) InventorySelectedCommands() []list.Item {
	return []list.Item{
		command{
			name:        "Check In Item",
			description: "Checks in the selected item",
		},
	}
}
