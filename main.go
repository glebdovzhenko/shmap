package main

import (
	"fmt"

	"github.com/glebdovzhenko/shmap/config"
	shdb "github.com/glebdovzhenko/shmap/database"
)

func main() {
    // loading population
	app_config := shcfg.GetConfig()
	fmt.Printf(
		"%s %d.%d.%d\n", app_config.Name,
		app_config.Version[0], app_config.Version[1], app_config.Version[2],
	)

    // loading database
    shdb.DefaultPopulate()
    cns, rs := shdb.GetTable("folders")
    fmt.Printf("%v\n%v\n", cns, rs)

}
