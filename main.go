package main

import (
	"fmt"

	"github.com/glebdovzhenko/shmap/config"
	"github.com/glebdovzhenko/shmap/database"
)

func main() {
	app_config := shcfg.GetConfig()
	fmt.Printf(
		"%s %d.%d.%d\n", app_config.Name,
		app_config.Version[0], app_config.Version[1], app_config.Version[2],
	)

    app_db := shdb.DefaultDB()
    defer app_db.Close()
    
    rows, err := app_db.Query("SELECT * FROM folders")
    if err != nil{
        panic(err)
    }
    fmt.Print("%v", rows)
}
