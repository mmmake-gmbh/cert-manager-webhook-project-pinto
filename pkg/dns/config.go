package dns

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"k8s.io/api/core/v1"
	extapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

const (
	defaultProvider          = "hexonet"
	defaultEnvironment       = "prod1"
	defaultPintoApiURL       = "https://api.stackit.domains"
	defaultOAuthTokenURL     = "https://auth.api.stackit.domains/connect/token"
	defaultOAuthClientID     = ""
	defaultOAuthClientSecret = ""
	defaultCredentialId      = ""
	ttlDNS                   = 60
	webhookName              = "pinto"
)

const (
	providerContextKey          = "provider"
	environmentContextKey       = "environment"
	pintoApiUrlContextKey       = "acme_url"
	oauthCTokenUrlContextKey    = "oauth_token_url"
	oauthClientIDContextKey     = "oauth_client_id"
	oauthClientSecretContextKey = "oauth_client_secret"
	oauthScopesContextKey       = "oauth_scopes"
	credentialIdContextKey      = "credential_id"
)

const (
	pintoProviderEnvName     = "PINTO_PROVIDER"
	oauthClientIDEnvName     = "PINTO_OAUTH_CLIENT_ID"
	oauthClientSecretEnvName = "PINTO_OAUTH_CLIENT_SECRET"
	pintoApiUrlEnvName       = "PINTO_API_URL"
	oauthTokenUrlEnvName     = "PINTO_OAUTH_TOKEN_URL"
	credentialIdEnvName      = "PINTO_CREDENTIAL_ID"
)

var (
	defaultOauthScopes = []string{
		"pinto_nexus",
		"fava_openapi_gateway",

		"fava_business_api",
		"fava_credentials_api",
		"pinto_citadel",
	}
)

type Config struct {
	savedContext context.Context
}

// ProviderConfig represents the config used for pinto DNS
type ProviderConfig struct {
	AccessKey     *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty"`
	SecretKey     *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty"`
	PintoProvider string                `json:"pintoProvider,omitempty"`
	PintoApiUrl   string                `json:"pintoApiUrl,omitempty"`
	OauthTokenUrl string                `json:"oauthTokenUrl,omitempty"`
	CredentialsId string                `json:"credentialsId,omitempty"`
}

func (c *Config) getContext() context.Context {
	return c.savedContext
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource. The Provider Name is used here
func (c *Config) Name() string {
	return webhookName
}

// Environment is referencing the environment of the Pinto API. Defaults to the prod1 environment
func (c *Config) Environment() string {
	environment := c.getContext().Value(environmentContextKey)
	if environment == nil {
		environment = defaultEnvironment
	}
	return environment.(string)
}

func (c *Config) Provider() string {
	provider := c.getContext().Value(providerContextKey)
	if provider == nil {
		provider = defaultProvider
	}
	return provider.(string)
}

// PintoApiURL returns the URL to ACME instance. Defaults to Pinto Primary
func (c *Config) PintoApiURL() string {
	pintoApiUrl := c.getContext().Value(pintoApiUrlContextKey)
	if pintoApiUrl == nil {
		pintoApiUrl = defaultPintoApiURL
	}
	return pintoApiUrl.(string)
}

func (c *Config) OauthTokenURL() string {
	oauthTokenURL := c.getContext().Value(oauthCTokenUrlContextKey)
	if oauthTokenURL == nil {
		oauthTokenURL = defaultOAuthTokenURL
	}
	return oauthTokenURL.(string)
}

func (c *Config) OauthClientID() string {
	oauthClientId := c.getContext().Value(oauthClientIDContextKey)
	if oauthClientId == nil {
		oauthClientId = defaultOAuthClientID
	}
	return oauthClientId.(string)
}

func (c *Config) OauthClientSecret() string {
	oauthClientSecret := c.getContext().Value(oauthClientSecretContextKey)
	if oauthClientSecret == nil {
		oauthClientSecret = defaultOAuthClientSecret
	}
	return oauthClientSecret.(string)
}

func (c *Config) OauthClientScopes() []string {
	oauthScopes := c.getContext().Value(oauthScopesContextKey)
	if oauthScopes == nil {
		return defaultOauthScopes
	}
	return oauthScopes.([]string)
}

func (c *Config) CredentialsId() string {
	credentialId := c.getContext().Value(credentialIdContextKey)
	if credentialId == nil {
		return defaultCredentialId
	}
	return credentialId.(string)
}

func (c *Config) init(k8Client kubernetes.Interface, ch *v1alpha1.ChallengeRequest) error {
	initialContext := context.Background()

	config, err := loadConfig(ch.Config)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	enrichedContext := initialContext

	// evaluate access and secret key
	accessKey := os.Getenv(oauthClientIDEnvName)
	secretKey := os.Getenv(oauthClientSecretEnvName)

	if config.AccessKey != nil && config.SecretKey != nil {
		accessKeySecret, err := k8Client.CoreV1().Secrets(ch.ResourceNamespace).Get(context.Background(), config.AccessKey.Name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("could not get secret %s: %w", config.AccessKey.Name, err)
		}
		secretKeySecret, err := k8Client.CoreV1().Secrets(ch.ResourceNamespace).Get(context.Background(), config.SecretKey.Name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("could not get secret %s: %w", config.SecretKey.Name, err)
		}

		accessKeyData, ok := accessKeySecret.Data[config.AccessKey.Key]
		if !ok {
			return fmt.Errorf("could not get key %s in secret %s", config.AccessKey.Key, config.AccessKey.Name)
		}

		secretKeyData, ok := secretKeySecret.Data[config.SecretKey.Key]
		if !ok {
			return fmt.Errorf("could not get key %s in secret %s", config.SecretKey.Key, config.SecretKey.Name)
		}

		accessKey = string(accessKeyData)
		secretKey = string(secretKeyData)
	}
	enrichedContext = context.WithValue(enrichedContext, oauthClientIDContextKey, accessKey)
	enrichedContext = context.WithValue(enrichedContext, oauthClientSecretContextKey, secretKey)

	// Pinto provider
	pintoProvider := defaultProvider
	pintoProviderEnvironment := os.Getenv(pintoProviderEnvName)
	if pintoProviderEnvironment != "" {
		pintoProvider = pintoProviderEnvironment
	}
	if config.PintoProvider != "" {
		pintoProvider = config.PintoProvider
	}
	enrichedContext = context.WithValue(enrichedContext, providerContextKey, pintoProvider)

	// evaluate API url
	pintoApiUrl := defaultPintoApiURL

	pintoApiUrlEnvironment := os.Getenv(pintoApiUrlEnvName)
	if pintoApiUrlEnvironment != "" {
		pintoApiUrl = pintoApiUrlEnvironment
	}
	if config.PintoApiUrl != "" {
		pintoApiUrl = config.PintoApiUrl
	}
	enrichedContext = context.WithValue(enrichedContext, pintoApiUrlContextKey, pintoApiUrl)

	// evaluate oauth Token URL
	oauthTokenUrl := defaultOAuthTokenURL
	oauthTokenUrlEnvironment := os.Getenv(oauthTokenUrlEnvName)
	if oauthTokenUrlEnvironment != "" {
		oauthTokenUrl = oauthTokenUrlEnvironment
	}
	if config.OauthTokenUrl != "" {
		oauthTokenUrl = config.OauthTokenUrl
	}
	enrichedContext = context.WithValue(enrichedContext, oauthCTokenUrlContextKey, oauthTokenUrl)

	// evaluate credentialId
	credentialId := defaultCredentialId
	credentialIdEnvironment := os.Getenv(credentialIdEnvName)
	if oauthTokenUrlEnvironment != "" {
		credentialId = credentialIdEnvironment
	}
	if config.CredentialsId != "" {
		credentialId = config.CredentialsId
	}
	enrichedContext = context.WithValue(enrichedContext, credentialIdContextKey, credentialId)

	c.savedContext = enrichedContext
	return nil
}

// loadConfig is a small helper function that decodes JSON configuration into
// the typed config struct.
func loadConfig(cfgJSON *extapi.JSON) (ProviderConfig, error) {
	cfg := ProviderConfig{}
	// handle the 'base case' where no configuration has been provided
	if cfgJSON == nil {
		return cfg, nil
	}
	if err := json.Unmarshal(cfgJSON.Raw, &cfg); err != nil {
		return cfg, fmt.Errorf("error decoding solver config: %v", err)
	}

	return cfg, nil
}
