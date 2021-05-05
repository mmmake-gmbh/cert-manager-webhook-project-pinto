package dns

import (
	"context"
	"fmt"
	"k8s.io/api/core/v1"
)

const defaultProvider = "RRPproxy"

// ProviderConfig represents the config used for pinto DNS
type ProviderConfig struct {
	AccessKey *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty"`
	SecretKey *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty"`
}

func (p *ProviderSolver) getContext() context.Context {
	return context.Background()
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource. Defaulting to "RRPproxy"
func (p *ProviderSolver) Name() string {
	provider := p.getContext().Value("provider")
	if provider == nil {
		provider = defaultProvider
	}
	return fmt.Sprintf("%v", provider)
}
