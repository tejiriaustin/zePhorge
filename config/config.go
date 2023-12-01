package config

import (
	"github.com/tejiriaustin/pioneer/templating"
)

type Config struct {
	Port     string            `json:"port"`
	RepoName string            `json:"repoName"`
	Files    *templating.Files `json:"-"`
}

func New(files *templating.Files) *Config {
	return &Config{Files: files}
}
