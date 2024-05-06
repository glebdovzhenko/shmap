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
    //shdb.DefaultPopulate()
    tables_data := shdb.GetDBData()

    // prepping and launching TUI
    m := tui.InitTuiModel()
    m = tui.InitTuiModelTable(m, &(*tables_data)[1])

    if _, err := tea.NewProgram(*m).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
