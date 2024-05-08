package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	shcfg "github.com/glebdovzhenko/shmap/config"
)

func (m TuiModel) updateTable(msg tea.Msg) (TuiModel, tea.Cmd) {
	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func InitTuiModelTable(md *TuiModel) *TuiModel {
	app_cfg := shcfg.GetConfig()

	// getting the number of rows and columns in our table
	n_cols, _ := md.DBData.ColumnsLen(md.TableID)
	n_rows, _ := md.DBData.RowsLen(md.TableID)
    
    // defining width variables
	col_widths := make([]int, n_cols)
	for ii := 0; ii < n_cols; ii++ {
		col_widths[ii] = 0
	}

	// getting rows
	var (
		rows    []table.Row
		row_ptr *[]string
	)
	for ii := 0; ii < n_rows; ii++ {
		row_ptr, _ = md.DBData.RowPtr(md.TableID, ii)

		//
		for jj := 0; jj < n_cols; jj++ {
			if col_widths[jj] < len((*row_ptr)[jj]) {
				col_widths[jj] = len((*row_ptr)[jj])
			}
		}
		rows = append(rows, *row_ptr)
	}

	//Setting up column and total widths
	for ii := 0; ii < n_cols; ii++ {
		if col_widths[ii] > app_cfg.TUITable.MaxColWidth {
			col_widths[ii] = app_cfg.TUITable.MaxColWidth
		}
	}
    total_width := 0
	for ii := 0; ii < n_cols; ii++ {
		total_width += col_widths[ii] 
	}
    if total_width > app_cfg.TUITable.MaxTotalWidth {
        total_width = app_cfg.TUITable.MaxTotalWidth
    }


	// getting columns
	var (
		columns   []table.Column
		col_title string
	)
	for ii := 0; ii < n_cols; ii++ {
		col_title, _ = md.DBData.ColumnName(md.TableID, ii)
		columns = append(columns, table.Column{Title: col_title, Width: col_widths[ii]})
	}

	// constructing table
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
        //table.WithWidth(total_width),
	)

	// defining table style
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

	// setting the table in model & returning
	md.Table = t
	return md
}
