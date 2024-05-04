package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	shdb "github.com/glebdovzhenko/shmap/database"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type TuiModel struct {
	Table table.Model
}

func (m TuiModel) Init() tea.Cmd { return nil }

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m TuiModel) View() string {
	return baseStyle.Render(m.Table.View()) + "\n"
}

func InitTuiModel() *TuiModel {
	return &TuiModel{}
}

func InitTuiModelTable(md *TuiModel, tb * shdb.DBTableData) *TuiModel {

	var columns []table.Column
	for _, cn := range tb.ColumnNames {
		columns = append(columns, table.Column{Title: cn, Width: 10})
	}

	var rows []table.Row
	for ii, _ := range tb.Rows {
		rows = append(rows, tb.Rows[ii])
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	md.Table = t

	return md
}
