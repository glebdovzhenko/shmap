package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/glebdovzhenko/shmap/config"
	shdb "github.com/glebdovzhenko/shmap/database"
	"github.com/glebdovzhenko/shmap/tui"
	"os"
)

func main() {
	// loading config
	app_config := shcfg.GetConfig()
	fmt.Printf(
		"%s %d.%d.%d\n", app_config.Name,
		app_config.Version[0], app_config.Version[1], app_config.Version[2],
	)

	// loading database
	//shdb.DefaultPopulate()
	column_names, row_data := shdb.GetTable("folders")
	fmt.Printf("%v\n%v\n", column_names, row_data)

	// making tui
	var columns []table.Column
	for _, cn := range column_names {
		columns = append(columns, table.Column{Title: cn, Width: 10})
	}

	var rows []table.Row
	for ii, _ := range row_data {
		rows = append(rows, row_data[ii])
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

    m := tui.TuiModel{Table: t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
