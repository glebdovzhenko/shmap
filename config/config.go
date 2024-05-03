package config

import (
    "os"
    "github.com/pelletier/go-toml/v2"
)

type ShmapCfg struct {
	Version []int
	Name    string
}

func defaultConfig() *ShmapCfg {
	result := ShmapCfg{}
	result.Version = []int{0, 0, 0}
	result.Name = "DEFAULT"
	return &result
}

func GetConfig() (*ShmapCfg) {
    cfg := defaultConfig()

    doc, err := os.ReadFile("shmap.toml")
    if err != nil {
        return cfg
    }

	err = toml.Unmarshal([]byte(doc), cfg)
	if err != nil {
		return cfg
	}

    return cfg
}
