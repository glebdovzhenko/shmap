package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TuiModel struct {
	FocusOn FocusState
	Table   table.Model
	List    list.Model
    TextInput   textinput.Model
}

func (m TuiModel) Init() tea.Cmd { return textinput.Blink }

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// handling application-wide commands
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.FocusOn, _ = nextFocusState(m.FocusOn)
		case "shift+tab":
			m.FocusOn, _ = prevFocusState(m.FocusOn)
		}
	}

	// handling widget-wide updates
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch m.FocusOn {
	case FcOnTable:
		m, cmd = m.updateTable(msg)
		cmds = append(cmds, cmd)
	case FcOnList:
		m, cmd = m.updateList(msg)
		cmds = append(cmds, cmd)
    case FcOnCmdLine:
        m, cmd = m.updateTextInput(msg)
        cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m TuiModel) View() string {

	tb_rendered := tableStyle.Render(m.Table.View())
	lst_rendered := listStyle.Render(m.List.View())
	cmd_rendered := cmdStyle.Render(m.TextInput.View())

	switch m.FocusOn {
	case FcOnList:
		lst_rendered = listStyleActive.Render(m.List.View())
	case FcOnTable:
		tb_rendered = tableStyleActive.Render(m.Table.View())
	case FcOnCmdLine:
		cmd_rendered = cmdStyleActive.Render(m.TextInput.View())
	default:
		panic("")
	}

	result := lipgloss.JoinHorizontal(lipgloss.Top, lst_rendered, tb_rendered)
	result = lipgloss.JoinVertical(lipgloss.Center, cmd_rendered, result)

	return result

}

func InitTuiModel() *TuiModel {
	return &TuiModel{FocusOn: FcOnTable}
}
