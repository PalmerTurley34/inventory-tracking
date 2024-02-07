package main

import "github.com/charmbracelet/lipgloss"

const (
	green  = lipgloss.Color("#2dc937")
	yellow = lipgloss.Color("#e7b416")
	red    = lipgloss.Color("#cc3232")
	grey   = lipgloss.Color("#aaaaaa")
)

var (
	focusedListStyle   = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.RoundedBorder(), true).BorderForeground(green)
	unFocusedListStyle = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.RoundedBorder(), true).BorderForeground(grey)
	successHeaderStyle = lipgloss.NewStyle().Foreground(green)
	failureHeaderStyle = lipgloss.NewStyle().Foreground(red)
	loadingHeaderStyle = lipgloss.NewStyle().Foreground(yellow)
)
