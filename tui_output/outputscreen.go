package tui_output

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OutputScreenModel struct {
	CmdOutput []string
}

func InitOutputScreenModel() *OutputScreenModel {
	m := &OutputScreenModel{CmdOutput: []string{"", "", ""}}
	return m
}

func (m OutputScreenModel) Init() tea.Cmd { return nil }

func (m OutputScreenModel) Update(msg tea.Msg) (OutputScreenModel, tea.Cmd) { return m, nil }

func (m OutputScreenModel) View() string {
	result := lipgloss.JoinHorizontal(lipgloss.Center, m.CmdOutput...)
	return result
}
