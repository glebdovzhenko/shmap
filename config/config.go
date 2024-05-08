package shcfg

import (
	//"fmt"
	//"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type TableCfg struct {
	MaxColWidth   int
	MaxTotalWidth int
}

type ShmapCfg struct {
	ConfigPath string
	Version    []int  // 3-integer version
	Name       string // should be SHMAP or something
	DBPath     string // Location of the database file
	DBType     string // Is passed to sql.Open as driver
	TUITable   TableCfg
}

// I follow the neovim rule here:
// If $XDG_CONFIG_HOME is set, then it is a dir $XDG_CONFIG_HOME/shmap
// and by default $HOME/.config/shmap
func setupPath() string {
	// Checking XDG_CONFIG_HOME
	xdg_home := os.Getenv("XDG_CONFIG_HOME")
	if _, err := os.Stat(xdg_home); os.IsNotExist(err) {
		xdg_home = ""
	}

	// Getting user home
	user_home, err := os.UserHomeDir()
	if err != nil {
		user_home = ""
	}

	if (len(user_home) == 0) && (len(xdg_home) == 0) {
		panic("No place to put my files")
	}

	// choosing one or the other
	var shmap_path string
	if len(xdg_home) != 0 {
		shmap_path = filepath.Join(xdg_home, "shmap")
	} else {
		shmap_path = filepath.Join(user_home, ".config", "shmap")
	}

	err = os.MkdirAll(shmap_path, 0o770)
	if err != nil {
		panic(err)
	}

	return shmap_path
}

func defaultConfig() *ShmapCfg {
	result := ShmapCfg{}

	result.ConfigPath = setupPath()
	result.Version = []int{0, 0, 1}
	result.Name = "SHMAP!"
	result.DBType = "sqlite3"
	result.DBPath = filepath.Join(result.ConfigPath, "shmap.db")
	result.TUITable.MaxColWidth = 20
    result.TUITable.MaxTotalWidth = 100
	return &result
}

func GetConfig() *ShmapCfg {
	cfg := defaultConfig()

	config_path := filepath.Join(cfg.ConfigPath, "shmap.toml")
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		b, _ := toml.Marshal(cfg)
		os.WriteFile(config_path, b, 0o770)
		return cfg
	}

	doc, err := os.ReadFile(config_path)
	if err != nil {
		return cfg
	}

	err = toml.Unmarshal([]byte(doc), cfg)
	if err != nil {
		return cfg
	}

	return cfg
}
