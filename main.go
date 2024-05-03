package main

import (
	"fmt"
	shcfg "github.com/glebdovzhenko/shmap/config"
)

func main() {
	app_config := shcfg.Default()
	fmt.Printf(
        "%s %d.%d.%d\n", app_config.Name, 
        app_config.Version[0], app_config.Version[1], app_config.Version[2],
    )

}
