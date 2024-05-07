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
	tea.Printf(
		"%s %d.%d.%d\n", app_config.Name,
		app_config.Version[0], app_config.Version[1], app_config.Version[2],
	)
    
    // loading database
    tables_data := shdb.GetDBData()

    // prepping TUI
    m := tui.InitTuiModel()
    m = tui.InitTuiModelList(m, tables_data)
    m = tui.InitTuiModelTable(m, &(*tables_data)[1])
    m = tui.InitTuiModelTextInput(m)

    // launching TUI
    if _, err := tea.NewProgram(*m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
