package gcloud

import (
	"github.com/Matir/ctftk/parser"
	"github.com/Matir/ctftk/providers/hosting"
	"github.com/Matir/ctftk/providers/hosting/kubernetes"
)

type GCloudProvider struct {
	*kubernetes.KubernetesProvider
	RootConfig *parser.RootConfig
}

func NewGCloudProvider(cfg *parser.RootConfig) (*GCloudProvider, error) {
	kp, err := kubernetes.NewKubernetesProvider(cfg)
	if err != nil {
		return nil, err
	}
	return &GCloudProvider{
		KubernetesProvider: kp,
	}, nil
}

func init() {
	hosting.RegisterHostingProvider("gcloud", func(cfg *parser.RootConfig) (hosting.HostingProvider, error) {
		return NewGCloudProvider(cfg)
	})
}
