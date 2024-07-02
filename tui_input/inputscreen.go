package tui_input

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	shcfg "github.com/glebdovzhenko/shmap/config"
    shcmd "github.com/glebdovzhenko/shmap/tui_cmds"
)

var (
	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
            Height(10).
            Width(60)
	listStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("240")).
			Margin(1, 2).
            Height(20).
            Width(20)
	cmdStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
            Height(5).
            Width(80)

	tableStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
            Height(10).
            Width(60)
	listStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("69")).
			Margin(1, 2).
            Height(20).
            Width(20)
	cmdStyleActive = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69")).
            Height(5).
            Width(80)
)

func (m InputScreenModel) updateTable(msg tea.Msg) (InputScreenModel, tea.Cmd) {
	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}


func InitTuiModelTable(md *InputScreenModel) *InputScreenModel {
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


func (m InputScreenModel) updateList(msg tea.Msg) (InputScreenModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, tea.Batch(shcmd.EmitSwitchTableMsg(m.List.Index())))
		}
	}

	m.List, cmd = m.List.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}


type item struct {
	title, desc string
}


func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func InitTuiModelList(md *InputScreenModel) *InputScreenModel {

	items := make([]list.Item, len(*md.DBData))
	var (
		title          string
		n_rows, n_cols int
	)
	for ii, _ := range items {
		title, _ = md.DBData.Name(ii)
		n_cols, _ = md.DBData.ColumnsLen(ii)
		n_rows, _ = md.DBData.RowsLen(ii)
		items[ii] = list.Item(item{
			title: title,
			desc:  fmt.Sprintf("%d columns %d rows", n_cols, n_rows),
		})
	}

	md.List = list.New(items, list.NewDefaultDelegate(), 20, 20)
	md.List.Title = "Tables"
	md.List.SetShowStatusBar(false)
	md.List.SetFilteringEnabled(false)
	return md
}


func (m InputScreenModel) updateTextInput(msg tea.Msg) (InputScreenModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, tea.Batch(shcmd.EmitTextSubmitMsg(m.TextInput.Value())))
		}
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}


func InitTuiModelTextInput(md *InputScreenModel) *InputScreenModel {
	md.TextInput = textinput.New()
	md.TextInput.Placeholder = "echo {{name}}"
	md.TextInput.Focus()
	md.TextInput.CharLimit = 156
	md.TextInput.Width = 20
	return md
}

