package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m TuiModel) updateTable(msg tea.Msg) (TuiModel, tea.Cmd) {
	var cmd tea.Cmd
	//switch msg := msg.(type) {
	//case tea.KeyMsg:
	//    switch msg.String() {
	//    case "enter":
	//        return m, tea.Batch(
	//            tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
	//        )
	//    }
	//}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func InitTuiModelTable(md *TuiModel) *TuiModel {
    
    n_cols, _ := md.DBData.ColumnsLen(md.TableID)
    n_rows, _ :=md.DBData.RowsLen(md.TableID)

	var (
        columns []table.Column
        col_title string
    )
    for ii := 0; ii < n_cols; ii++ {
        col_title, _ = md.DBData.ColumnName(md.TableID, ii)
        columns = append(columns, table.Column{Title: col_title, Width: 10})
    }

	//for _, cn := range *(*(*md).DBData)[(*md).TableID].ColumnNames {
	//    columns = append(columns, table.Column{Title: cn, Width: 10})
	//}

	var (
        rows []table.Row
        row_ptr *[]string
    )

    for ii := 0; ii < n_rows; ii++ {
        row_ptr, _ = md.DBData.RowPtr(md.TableID, ii)
        rows = append(rows, *row_ptr)
    }
	//for ii, _ := range *(*(*md).DBData)[(*md).TableID].Rows {
    //    rows = append(rows, (*(*(*md).DBData)[(*md).TableID].Rows)[ii])
	//}

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
