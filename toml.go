package main

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	OutputDir  string `toml:"output_dir"`
	WorkingDir string `toml:"working_dir"`
	ZipUrl     string `toml:"zip_url"`
}

func (c *Config) Zipfile() string {
	return filepath.Join(c.WorkingDir, filepath.Base(c.ZipUrl))
}

func loadToml(file string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return &conf, err
	}

	return &conf, nil
}
