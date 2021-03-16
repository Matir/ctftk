package kubernetes

import (
	"github.com/Matir/ctftk/parser"
	"github.com/Matir/ctftk/providers/hosting"
)

type KubernetesProvider struct {
	RootConfig *parser.RootConfig
}

func NewKubernetesProvider(cfg *parser.RootConfig) (*KubernetesProvider, error) {
	return &KubernetesProvider{
		RootConfig: cfg,
	}, nil
}

func (kp *KubernetesProvider) String() string {
	return "KubernetesProvider"
}

func (kp *KubernetesProvider) GetChallengeStatus(*parser.ChallengeConfig) (hosting.HostedChallengeStatus, error) {
	return hosting.HostedChallengeStatus{}, hosting.NotImplementedError
}

func (kp *KubernetesProvider) UpdateChallenge(cfg *parser.ChallengeConfig) (hosting.HostedChallengeStatus, error) {
	return hosting.HostedChallengeStatus{}, hosting.NotImplementedError
}

func init() {
	hosting.RegisterHostingProvider("kubernetes", func(cfg *parser.RootConfig) (hosting.HostingProvider, error) {
		return NewKubernetesProvider(cfg)
	})
}
