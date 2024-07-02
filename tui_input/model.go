package tui_input

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    shdb "github.com/glebdovzhenko/shmap/database"
    shcmd "github.com/glebdovzhenko/shmap/tui_cmds"
	"log"
)


type FocusState uint

const (
	FcOnTable   FocusState = iota
	FcOnCmdLine FocusState = iota
	FcOnList    FocusState = iota
)

func (fs FocusState) next() (FocusState, error) {
	switch fs {
	case FcOnTable:
		return FcOnCmdLine, nil
	case FcOnCmdLine:
		return FcOnList, nil
	case FcOnList:
		return FcOnTable, nil
	default:
		return fs, fmt.Errorf("FocusState has unknown value %d", fs)
	}
}

func (fs FocusState) prev() (FocusState, error) {
	switch fs {
	case FcOnTable:
		return FcOnList, nil
	case FcOnCmdLine:
		return FcOnTable, nil
	case FcOnList:
		return FcOnCmdLine, nil
	default:
		return fs, fmt.Errorf("FocusState has unknown value %d", fs)
	}
}

type InputScreenModel struct {
	// data
	FocusOn FocusState
	DBData  *shdb.DBData
	TableID int

	//tea models
	Table     table.Model
	List      list.Model
	TextInput textinput.Model
}

func InitInputScreenModel(tables *shdb.DBData) *InputScreenModel {
	m := &InputScreenModel{
		FocusOn:   FcOnTable,
		DBData:    tables,
		TableID:   0,
	}
	m = InitTuiModelList(m)
	m = InitTuiModelTable(m)
	m = InitTuiModelTextInput(m)
	return m
}

func (m InputScreenModel) Update(msg tea.Msg) (InputScreenModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			log.Printf("TuiModel.Update received tab")
			m.FocusOn, _ = m.FocusOn.next()
		case "shift+tab":
			log.Printf("TuiModel.Update received shift+tab")
			m.FocusOn, _ = m.FocusOn.prev()
		}
	case shcmd.SwitchTableMsg:
		log.Printf("TuiModel.Update received SwitchTableMsg: %d", msg)

		m.TableID = int(msg)
		InitTuiModelTable(&m)
        var cmd tea.Cmd
        m, cmd = m.updateTable(msg)
		return m, cmd

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


func (m InputScreenModel) View() string {

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


