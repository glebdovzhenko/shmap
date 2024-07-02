package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type SwitchTableMsg int

func emitSwitchTableMsg(tb_id int) tea.Cmd {
	return func() tea.Msg {
		return SwitchTableMsg(tb_id)
	}
}

type TextSubmitMsg string

func emitTextSubmitMsg(data string) tea.Cmd {
	return func() tea.Msg {
		return TextSubmitMsg(data)
	}
}

type WorkerResultMsg string

func runWorkerCmd(input string) (output tea.Cmd) {
	return func() tea.Msg {
		time.Sleep(100 * time.Millisecond)
		return WorkerResultMsg(input)
	}
}
