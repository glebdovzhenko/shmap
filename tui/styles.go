package tui

import "github.com/charmbracelet/lipgloss"

var (
	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
            Height(10).
            Width(60)
	listStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("240")).
			Margin(1, 2).
            Height(20).
            Width(20)
	cmdStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
            Height(5).
            Width(80)

	tableStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
            Height(10).
            Width(60)
	listStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("69")).
			Margin(1, 2).
            Height(20).
            Width(20)
	cmdStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
            Height(5).
            Width(80)
)
