package config

import (
	"github.com/BurntSushi/toml"
)

func LoadConfigFromFile(filePath string, destination interface{}) error {
	_, err := toml.DecodeFile(filePath, destination)
	return err
}
