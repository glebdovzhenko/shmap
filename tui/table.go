package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	shdb "github.com/glebdovzhenko/shmap/database"
)

func (m TuiModel) updateTable(msg tea.Msg) (TuiModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func InitTuiModelTable(md *TuiModel, tb *shdb.DBTableData) *TuiModel {

	var columns []table.Column
	for _, cn := range *tb.ColumnNames {
		columns = append(columns, table.Column{Title: cn, Width: 10})
	}

	var rows []table.Row
	for ii, _ := range *tb.Rows {
		rows = append(rows, (*tb.Rows)[ii])
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		//table.WithHeight(20),
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
