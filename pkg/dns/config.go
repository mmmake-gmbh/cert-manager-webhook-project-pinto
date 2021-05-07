package dns

import (
	"context"
	"gitlab.com/whizus/gopinto"
	"k8s.io/api/core/v1"
	"strings"
)

const (
	defaultProvider      = "RRPproxy"
	defaultEnvironment   = "prod1"
	defaultAcmeURL       = "" // TODO change default value
	defaultOAuthTokenURL = "" // TODO change default value
	ttlDNS               = 60
)

var (
	defaultOauthScopes = []string{
		"citadel",
		"acmegateway",
	}
)

type Config struct {
	savedContext context.Context
}

// ProviderConfig represents the config used for pinto DNS
type ProviderConfig struct {
	AccessKey *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty"`
	SecretKey *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty"`
}

func (c *Config) getContext() context.Context {
	return c.savedContext
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource. Defaulting to "RRPproxy"
func (c *Config) Name() string {
	provider := c.getContext().Value("provider")
	if provider == nil {
		provider = defaultProvider
	}
	return provider.(string)
}

// Environment is referencing the environment of the Pinto API. Defaults to the prod1 environment
func (c *Config) Environment() gopinto.NullableString {
	environment := c.getContext().Value("environment")
	if environment == nil {
		environment = defaultEnvironment
	}
	resultString := environment.(string)
	result := new(gopinto.NullableString)
	result.Set(&resultString)
	return *result
}

// ACMEServerURL returns the URL to ACME instance. Defaults to Pinto Primary
func (c *Config) ACMEServerURL() string {
	acmeURL := c.getContext().Value("acme_url")
	if acmeURL == nil {
		acmeURL = defaultAcmeURL
	}
	return acmeURL.(string)
}

func (c *Config) OauthTokenURL() string {
	oauthTokenURL := c.getContext().Value("oauth_token_url")
	if oauthTokenURL == nil {
		oauthTokenURL = defaultOAuthTokenURL
	}
	return oauthTokenURL.(string)
}

func (c *Config) OauthClientID() string {
	// TODO implement
	return ""
}

func (c *Config) OauthClientSecret() string {
	// TODO implement
	return ""
}

func (c *Config) OauthClientScopes() []string {
	oauthScopesString := c.getContext().Value("oauth_scopes")
	if oauthScopesString == nil {
		return defaultOauthScopes
	}
	scopes := strings.Split(oauthScopesString.(string), ",")
	var massagedScopes []string
	for _, scope := range scopes {
		massagedScopes = append(massagedScopes, strings.Trim(scope, " "))
	}
	return massagedScopes
}
