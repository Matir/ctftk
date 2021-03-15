package parser

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

type HostingConfig struct {
	Name string `yaml:"name"`
}

type ScoreboardConfig struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type RootConfig struct {
	Hosting    HostingConfig    `yaml:"hosting"`
	Scoreboard ScoreboardConfig `yaml:"scoreboard"`
	Version    string           `yaml:"version"`
}

type ConfigError struct {
	ErrMsg string
}

func NewRootConfig() *RootConfig {
	return &RootConfig{}
}

func ReadRootConfig(r io.Reader) (*RootConfig, error) {
	yamlDec := yaml.NewDecoder(r)
	rv := NewRootConfig()
	if err := yamlDec.Decode(rv); err != nil {
		return nil, NewConfigError(err.Error())
	}
	if err := rv.Valid(); err != nil {
		return nil, err
	}
	return rv, nil
}

// Returns nil if the config is valid, otherwise returns an error describing
// why it is invalid.
func (rc *RootConfig) Valid() error {
	if rc.Version == "" {
		return NewConfigError("Version is required.")
	}
	if rc.Hosting.Name == "" {
		return NewConfigError("Host name is required.")
	}
	return nil
}

func (ce ConfigError) Error() string {
	return ce.ErrMsg
}

func NewConfigError(msg string) ConfigError {
	return ConfigError{fmt.Sprintf("ConfigError: %s", msg)}
}
