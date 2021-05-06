package dns

import (
	"context"
	"gitlab.com/whizus/gopinto"
	"k8s.io/api/core/v1"
)

const defaultProvider = "RRPproxy"
const defaultEnvironment = "prod1"
const ttlDNS = 60

var savedContext = context.Background()

// ProviderConfig represents the config used for pinto DNS
type ProviderConfig struct {
	AccessKey *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty"`
	SecretKey *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty"`
}

func (p *ProviderSolver) getContext() context.Context {
	return savedContext
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource. Defaulting to "RRPproxy"
func (p *ProviderSolver) Name() string {
	provider := p.getContext().Value("provider")
	if provider == nil {
		provider = defaultProvider
	}
	return provider.(string)
}

// Environment is referencing the environment of the Pinto API. Defaults to the prod1 environment
func (p *ProviderSolver) Environment() gopinto.NullableString {
	environment := p.getContext().Value("environment")
	if environment == nil {
		environment = defaultEnvironment
	}
	resultString := environment.(string)
	result := new(gopinto.NullableString)
	result.Set(&resultString)
	return *result
}
