package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	//shdb "github.com/glebdovzhenko/shmap/database"
)

func (m TuiModel) updateTextInput(msg tea.Msg) (TuiModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, tea.Batch(emitTextSubmitMsg(m.TextInput.Value())))
		}
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func InitTuiModelTextInput(md *TuiModel) *TuiModel {
	md.TextInput = textinput.New()
	md.TextInput.Placeholder = "Pikachu"
	md.TextInput.Focus()
	md.TextInput.CharLimit = 156
	md.TextInput.Width = 20
	return md
}
