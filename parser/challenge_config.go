package parser

import (
	"io"
	"strings"

	"gopkg.in/yaml.v2"
)

type ChallengeType int

const (
	ChallengeTypeHTTP ChallengeType = iota
	ChallengeTypeTCP
	ChallengeTypeOffline
)

var (
	ChallengeTypes = map[ChallengeType]string{
		ChallengeTypeHTTP:    "http",
		ChallengeTypeTCP:     "tcp",
		ChallengeTypeOffline: "offline",
	}
)

type ChallengeConfig struct {
	ChallID          string           `yaml:"-"`
	Name             string           `yaml:"name"`
	Version          string           `yaml:"version"`
	Description      string           `yaml:"description"`
	TypeString       string           `yaml:"type"`
	ChallengeType    ChallengeType    `yaml:"-"`
	Author           string           `yaml:"author"`
	Flag             string           `yaml:"flag"`
	Flags            []FlagConfig     `yaml:"flags"`
	Points           int              `yaml:"points"`
	ContainerConfig  ContainerConfig  `yaml:"container"`
	DeploymentConfig DeploymentConfig `yaml:"deployment"`
	Notes            string           `yaml:"notes"`
}

type FlagConfig struct {
	Name      string `yaml:"name"`
	Flag      string `yaml:"flag"`
	Points    int    `yaml:"points"`
	Validator string `yaml:"validator"`
}

type ContainerConfig struct {
	BuildCommand    string `yaml:"build_command"`
	PrebuildCommand string `yaml:"prebuild_command"`
}

type DeploymentConfig struct {
	Setuid bool `yaml:"setuid"`
	Ptrace bool `yaml:"ptrace"`
	Port   int  `yaml:"port"`
}

func NewChallengeConfig(id string) *ChallengeConfig {
	return &ChallengeConfig{
		ChallID: id,
		Version: "latest",
	}
}

func ReadChallengeConfig(r io.Reader, id string) (*ChallengeConfig, error) {
	rv := NewChallengeConfig(id)
	yamlDec := yaml.NewDecoder(r)
	yamlDec.SetStrict(true)
	if err := yamlDec.Decode(rv); err != nil {
		return nil, NewConfigError(err.Error())
	}
	if err := rv.SetChallengeTypeFromString(rv.TypeString); err != nil {
		return nil, err
	}
	if err := rv.Valid(); err != nil {
		return nil, err
	}
	return rv, nil
}

// Returns nil if valid, otherwise returns an error describing the validation
// issues
func (cc *ChallengeConfig) Valid() error {
	if cc.Name == "" {
		return NewConfigError("Name is required")
	}
	if cc.Flag != "" && len(cc.Flags) > 0 {
		return NewConfigError("Must specify one of `flag` or `flags`, not both.")
	}
	return nil
}

func (cc *ChallengeConfig) SetChallengeType(t ChallengeType) {
	cc.ChallengeType = t
	cc.TypeString = ChallengeTypes[t]
}

func (cc *ChallengeConfig) SetChallengeTypeFromString(t string) error {
	t = strings.ToLower(t)
	for k, v := range ChallengeTypes {
		if v == t {
			cc.ChallengeType = k
			cc.TypeString = t
			return nil
		}
	}
	return NewConfigError("Invalid challenge type: %s", t)
}
