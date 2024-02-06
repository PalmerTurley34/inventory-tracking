package main

import (
	"fmt"
	"os"

	"github.com/PalmerTurley34/inventory-tracking/internal/backend"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	go backend.StartBackendServer()

	m := newModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
