package config

import (
	"bufio"
	"github.com/BurntSushi/toml"
	"github.com/manlycode/go-to-work/service"
	"os"
)

// Config defines a data structure that contains configuration for the application
type Config struct {
	Services map[string]*service.Service
}

// ConfigFile defines a file to store the configs
type ConfigFile struct {
	Path string
}

const defaultPath = ".go-to-work"

// NewConfigFile builds a ConfigFile in the current directory w/ the path '.go-to-work'
func NewConfigFile() (file *ConfigFile) {
	return &ConfigFile{Path: defaultPath}
}

// Save writes the configuration
func (c *Config) Save() error {
	f, err := os.Create(defaultPath)
	writer := bufio.NewWriter(f)

	encoder := toml.NewEncoder(writer)
	encoder.Encode(c)
	f.Close()

	return err
}
