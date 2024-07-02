package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	shdb "github.com/glebdovzhenko/shmap/database"
)

type ScreenState uint

const (
	InputScreen  ScreenState = iota
	OutputScreen ScreenState = iota
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

type TuiModel struct {
	// data
	Screen  ScreenState
	FocusOn FocusState
	DBData  *shdb.DBData
	TableID int
    CmdOutput []string

	//tea models
	Table     table.Model
	List      list.Model
	TextInput textinput.Model
}

func InitTuiModel(tables *shdb.DBData) *TuiModel {
	m := &TuiModel{
        Screen: InputScreen, 
        FocusOn: FcOnTable, 
        DBData: tables, 
        TableID: 0, 
        CmdOutput: []string{"", "", ""},
    }
	m = InitTuiModelList(m)
	m = InitTuiModelTable(m)
	m = InitTuiModelTextInput(m)
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
		}
	}

	switch m.Screen {
	case InputScreen:
		return m.updateInputScreen(msg)
	case OutputScreen:
		return m.updateOutputScreen(msg)
	default:
		panic("")
	}
}

func (m TuiModel) View() string {
	switch m.Screen {
	case InputScreen:
		return m.viewInputScreen()
	case OutputScreen:
		return m.viewOutputScreen()
	default:
		panic("")
	}
}
