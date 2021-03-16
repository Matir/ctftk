package hosting

import (
	"fmt"

	"github.com/Matir/ctftk/parser"
)

type HostingProvider interface {
	GetChallengeStatus(*parser.ChallengeConfig) (HostedChallengeStatus, error)
	UpdateChallenge(*parser.ChallengeConfig) (HostedChallengeStatus, error)
	String() string
	// TODO: Create/update DNS records, IP allocations
}

type ChallengeState int

const (
	ChallengeNotFound ChallengeState = iota
	ChallengeRunning
	ChallengeFailed
)

type HostedChallengeStatus struct {
	ChallengeState ChallengeState
	Replicas       int
}

// Registry of Hosting Providers
type HostingProviderConstructor func(*parser.RootConfig) (HostingProvider, error)

var hostingProviderRegistry = make(map[string]HostingProviderConstructor)

// Instantiate the hosting provider based on the config
func CreateHostingProvider(cfg *parser.RootConfig) (HostingProvider, error) {
	if constructor, ok := hostingProviderRegistry[cfg.Hosting.Name]; !ok {
		return nil, NewHostingError("Unknown hosting provider: %s", cfg.Hosting.Name)
	} else {
		return constructor(cfg)
	}
}

// Register a hosting provider.
func RegisterHostingProvider(name string, f HostingProviderConstructor) {
	hostingProviderRegistry[name] = f
}

// Special errors for hosting
type HostingError string

func (he HostingError) Error() string {
	return string(he)
}

func NewHostingError(e string, args ...interface{}) error {
	return HostingError(fmt.Sprintf(e, args...))
}

const (
	NotImplementedError = HostingError("Not implemented")
)
