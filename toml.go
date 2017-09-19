package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	WorkingDir string `toml:"working_dir"`
	ZipUrl     string `toml:"zip_url"`
}

func loadToml() *Config {
	var conf Config
	if _, err := toml.DecodeFile("", &conf); err != nil {
		panic(err)
	}

	return &conf
}
