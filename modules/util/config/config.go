package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessary for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
}

// Database holds data necessary for database configuration
type Database struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	Database string `yaml:"database,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	LogQueries bool `yaml:"log_queries,omitempty"`
	Timeout    int  `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessary for server configuration
type Server struct {
	Mode string `yaml:"mode,omitempty"`
	Port string `yaml:"port,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}
