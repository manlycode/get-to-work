package config

import (
	"bufio"
	"github.com/BurntSushi/toml"
	"os"
)

type service interface {
	Name() string
}

// Config defines a data structure that contains configuration for the application
type Config struct {
	Services map[string]*service
}

// NewConfig creates a new Config instance
func NewConfig() (c *config) {
	return &Config{Services: make(map[string]*service)}
}

// ConfigFile defines a file to store the configs
type ConfigFile struct {
	Path string
}

func (c *Config) AddService(s service) {
	services := &c.Services
	services[s.Name()] = s
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
