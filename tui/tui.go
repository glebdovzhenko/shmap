package tui

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	shdb "github.com/glebdovzhenko/shmap/database"
	shcmd "github.com/glebdovzhenko/shmap/tui_cmds"
	shin "github.com/glebdovzhenko/shmap/tui_input"
    shou "github.com/glebdovzhenko/shmap/tui_output"

)

type ScreenState uint

const (
	InputScreen  ScreenState = iota
	OutputScreen ScreenState = iota
)

type TuiModel struct {
	// data
	Screen       ScreenState
	InputScreen  shin.InputScreenModel
	OutputScreen shou.OutputScreenModel
}

func InitTuiModel(tables *shdb.DBData) *TuiModel {
	m := &TuiModel{
		Screen:       InputScreen,
		InputScreen:  *shin.InitInputScreenModel(tables),
		OutputScreen: *shou.InitOutputScreenModel(),
	}
	return m
}

func (m TuiModel) Init() tea.Cmd { return textinput.Blink }

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// handling application-wide commands
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			log.Printf("TuiModel.Update received ctrl+c")
			return m, tea.Quit
		case "enter":
			if m.Screen == OutputScreen {
				m.Screen = InputScreen
				return m, nil
			}
		}
	case shcmd.TextSubmitMsg:
		log.Printf("TuiModel.Update received TextSubmitMsg: \"%s\"", msg)
        
		return m, shcmd.RunWorkerCmd(string(msg))
	case shcmd.WorkerResultMsg:
		m.OutputScreen.CmdOutput[0] = string(msg)
		m.Screen = OutputScreen
		return m, nil
	}

	switch m.Screen {
	case InputScreen:
		var cmd tea.Cmd
		m.InputScreen, cmd = m.InputScreen.Update(msg)
		return m, cmd
	case OutputScreen:
		var cmd tea.Cmd
		m.OutputScreen, cmd = m.OutputScreen.Update(msg)
		return m, cmd
	default:
		panic("")
	}
}

func (m TuiModel) View() string {
	switch m.Screen {
	case InputScreen:
		return m.InputScreen.View()
	case OutputScreen:
		return m.OutputScreen.View()
	default:
		panic("")
	}
}
