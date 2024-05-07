package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type SwitchTableMsg int

func emitSwitchTableMsg(tb_id int) tea.Cmd {
	return func() tea.Msg {
		return SwitchTableMsg(tb_id)
	}
}
