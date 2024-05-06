package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// in the future these will be TuiModel.FocusOn values
//const (
//    FcOnTable   = iota
//    FcOnCmdLine = iota
//)

type TuiModel struct {
	FocusOn int
	Table   table.Model
}

func (m TuiModel) Init() tea.Cmd { return nil }

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// handling application-wide commands
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	// handling widget-wide updates
	var cmds []tea.Cmd
	var cmd tea.Cmd
	// updating table
	m, cmd = m.updateTable(msg)
	cmds = append(cmds, cmd)
    // updating other widgets...
    // ...
	return m, tea.Batch(cmds...)
}

func (m TuiModel) View() string {
	return baseStyle.Render(m.Table.View()) + "\n"
}

func InitTuiModel() *TuiModel {
	return &TuiModel{}
}
