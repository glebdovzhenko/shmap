package shcfg

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

type ShmapCfg struct {
	Version []int  // 3-integer version
	Name    string // should be SHMAP or something
	DBLoc   string // Location of the database file
	DBType  string // Is passed to sql.Open as driver
}

func defaultConfig() *ShmapCfg {
	result := ShmapCfg{}
	result.Version = []int{0, 0, 0}
	result.Name = "DEFAULT"
	return &result
}

func GetConfig() *ShmapCfg {
	cfg := defaultConfig()

	doc, err := os.ReadFile("./.appdata/shmap.toml")
	if err != nil {
		return cfg
	}

	err = toml.Unmarshal([]byte(doc), cfg)
	if err != nil {
		return cfg
	}

	return cfg
}
