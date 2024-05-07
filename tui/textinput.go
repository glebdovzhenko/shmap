package tui

import (
    tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"

    //shdb "github.com/glebdovzhenko/shmap/database"

)
func (m TuiModel) updateTextInput(msg tea.Msg) (TuiModel, tea.Cmd) {
    var cmd tea.Cmd
    m.TextInput, cmd = m.TextInput.Update(msg)	
	return m, cmd
}



func InitTuiModelTextInput(md *TuiModel) *TuiModel {
    md.TextInput = textinput.New()
	md.TextInput.Placeholder = "Pikachu"
	md.TextInput.Focus()
	md.TextInput.CharLimit = 156
	md.TextInput.Width = 20
	return md
}
