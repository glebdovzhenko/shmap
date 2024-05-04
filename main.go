package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/glebdovzhenko/shmap/config"
	shdb "github.com/glebdovzhenko/shmap/database"
	"github.com/glebdovzhenko/shmap/tui"
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

	// prepping and launching TUI
	m := tui.InitTuiModel()
	m = tui.InitTuiModelTable(m, &column_names, &row_data)

	if _, err := tea.NewProgram(*m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
