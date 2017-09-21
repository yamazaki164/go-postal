package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	OutputDir  string `toml:"output_dir"`
	WorkingDir string `toml:"working_dir"`
	ZipUrl     string `toml:"zip_url"`
}

func (c *Config) ZipName() string {
	return filepath.Base(c.ZipUrl)
}

func (c *Config) ZipFile() string {
	return filepath.Join(c.WorkingDir, c.ZipName())
}

func (c *Config) isValidDir(dir string) bool {
	st, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func (c *Config) IsValidConfig() bool {
	return c.isValidDir(c.OutputDir) && c.isValidDir(c.WorkingDir)
}

func loadToml(file string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(file, &conf); err != nil {
		return &conf, err
	}

	return &conf, nil
}
