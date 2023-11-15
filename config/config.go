package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	DummyConf string `toml:"dummy_conf"`
}

func init() {
	_ = godotenv.Load("service.toml")
}
