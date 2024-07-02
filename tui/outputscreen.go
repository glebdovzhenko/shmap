package tui
import 	(
    "github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"

)

type OutputScreenModel struct {
	CmdOutput []string
}

func InitOutputScreenModel() *OutputScreenModel {
    m:= &OutputScreenModel{CmdOutput: []string{"", "", ""}}
    return m
}

func (m TuiModel) updateOutputScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String(){
        case "enter":
            m.Screen = InputScreen
            return m, nil
        }
    }
	return m, nil
}

func (m TuiModel) viewOutputScreen() string {
    result := lipgloss.JoinHorizontal(lipgloss.Center, m.OutputScreen.CmdOutput...)
	return result
}

