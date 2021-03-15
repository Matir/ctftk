package hosting

import (
	"errors"

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

type HostingError string

func (he HostingError) Error() string {
	return string(he)
}

const (
	NotImplementedError = HostingError("Not implemented")
)
