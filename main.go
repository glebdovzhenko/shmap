package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/glebdovzhenko/shmap/config"
	shdb "github.com/glebdovzhenko/shmap/database"
	"github.com/glebdovzhenko/shmap/tui"
)

func main() {
    //setting up logging
    if len(os.Getenv("DEBUG")) > 0 {
        f, err := tea.LogToFile("debug.log", "debug")
        if err != nil {
            fmt.Println("fatal:", err)
            os.Exit(1)
        }
        defer f.Close()
    }
    
    log.Printf("Starting up SHMAP...")

	// loading config
	app_config := shcfg.GetConfig()
	tea.Printf(
		"%s %d.%d.%d\n", app_config.Name,
		app_config.Version[0], app_config.Version[1], app_config.Version[2],
	)

	// loading database
	tables_data := shdb.GetDBData()

	// prepping TUI
	m := tui.InitTuiModel(tables_data)

	// launching TUI
	if _, err := tea.NewProgram(*m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
