package parser

import (
	"io"

	"gopkg.in/yaml.v2"
)

// Currently this is a union of all possible fields.  In the future this may be
// split to per-provider structs.
type HostingConfig struct {
	Name           string `yaml:"name"`
	ProjectID      string `yaml:"project_id"`
	Region         string `yaml:"region"`
	Zone           string `yaml:"zone"`
	ServiceAccount string `yaml:"service_account"`
	DNSZoneName    string `yaml:"challenge_dns_zone"`
}

type ScoreboardConfig struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type RootConfig struct {
	Hosting         HostingConfig    `yaml:"hosting"`
	Scoreboard      ScoreboardConfig `yaml:"scoreboard"`
	Version         string           `yaml:"version"`
	ChallengeDomain string           `yaml:"challenge_domain"`
	DefaultReplicas uint16           `yaml:"default_replicas"`
	MaxReplicas     uint16           `yaml:"max_replicas"`
	Tags            []TagConfig      `yaml:"tags"`
}

type TagConfig struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Aliases     []string `yaml:"aliases"`
}

const (
	DefaultHostingProvider    = "gcloud"
	DefaultScoreboardProvider = "ctfscoreboard"
	DefaultGCloudRegion       = "us-west1"
	DefaultGCloudZone         = "us-west1-a"
)

func NewRootConfig() *RootConfig {
	return &RootConfig{
		Hosting: HostingConfig{
			Name:   DefaultHostingProvider,
			Region: DefaultGCloudRegion,
			Zone:   DefaultGCloudZone,
		},
		Scoreboard: ScoreboardConfig{
			Name: DefaultScoreboardProvider,
		},
		DefaultReplicas: 1,
	}
}

func ReadRootConfig(r io.Reader) (*RootConfig, error) {
	yamlDec := yaml.NewDecoder(r)
	yamlDec.SetStrict(true)
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
