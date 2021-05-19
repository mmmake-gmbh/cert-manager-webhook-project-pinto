package dns

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"gitlab.com/whizus/gopinto"
	"k8s.io/api/core/v1"
	extapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

const (
	defaultProvider          = "RRPproxy"
	defaultEnvironment       = "prod1"
	defaultPintoApiURL       = "https://pinto.irgendwo.co"
	defaultOAuthTokenURL     = "https://auth.pinto.irgendwo.co/connect/token"
	defaultOAuthClientID     = ""
	defaultOAuthClientSecret = ""
	ttlDNS                   = 60
)

const (
	providerContextKey          = "provider"
	environmentContextKey       = "environment"
	pintoApiUrlContextKey       = "acme_url"
	oauthCTokenUrlContextKey    = "oauth_token_url"
	oauthClientIDContextKey     = "oauth_client_id"
	oauthClientSecretContextKey = "oauth_client_secret"
	oauthScopesContextKey       = "oauth_scopes"
)

const (
	oauthClientIDEnvName     = "PINTO_OAUTH_CLIENT_ID"
	oauthClientSecretEnvName = "PINTO_OAUTH_CLIENT_SECRET"
	pintoApiUrlEnvName       = "PINTO_API_URL"
	oauthTokenUrlEnvName     = "PINTO_OAUTH_TOKEN_URL"
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
	AccessKey     *v1.SecretKeySelector `json:"accessKeySecretRef,omitempty"`
	SecretKey     *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty"`
	PintoApiUrl   *v1.SecretKeySelector `json:"pintoApiUrlSecretRef,omitempty"`
	OauthTokenUrl *v1.SecretKeySelector `json:"oauthTokenUrlSecretRef,omitempty"`
}

func (c *Config) getContext() context.Context {
	return c.savedContext
}

// Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource. Defaulting to "RRPproxy"
func (c *Config) Name() string {
	return c.Provider()
}

// Environment is referencing the environment of the Pinto API. Defaults to the prod1 environment
func (c *Config) Environment() gopinto.NullableString {
	environment := c.getContext().Value(environmentContextKey)
	if environment == nil {
		environment = defaultEnvironment
	}
	resultString := environment.(string)
	result := new(gopinto.NullableString)
	result.Set(&resultString)
	return *result
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

	// evaluate API url
	pintoApiUrl := os.Getenv(pintoApiUrlEnvName)
	if config.PintoApiUrl != nil {
		pintoApiUrlSecret, err := k8Client.CoreV1().Secrets(ch.ResourceNamespace).Get(context.Background(), config.PintoApiUrl.Name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("could not get secret %s: %w", config.PintoApiUrl.Name, err)
		}
		pintoApiUrlData, ok := pintoApiUrlSecret.Data[config.PintoApiUrl.Key]
		if !ok {
			return fmt.Errorf("could not get key %s in secret %s", config.PintoApiUrl.Key, config.PintoApiUrl.Name)
		}
		pintoApiUrl = string(pintoApiUrlData)
	}
	enrichedContext = context.WithValue(enrichedContext, pintoApiUrlContextKey, pintoApiUrl)

	// evaluate oauth Token URL
	oauthTokenUrl := os.Getenv(oauthTokenUrlEnvName)
	if config.OauthTokenUrl != nil {
		oauthTokenUrlSecret, err := k8Client.CoreV1().Secrets(ch.ResourceNamespace).Get(context.Background(), config.OauthTokenUrl.Name, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("could not get secret %s: %w", config.OauthTokenUrl.Name, err)
		}
		oauthTokenUrlData, ok := oauthTokenUrlSecret.Data[config.OauthTokenUrl.Key]
		if !ok {
			return fmt.Errorf("could not get key %s in secret %s", config.OauthTokenUrl.Key, config.OauthTokenUrl.Name)
		}
		oauthTokenUrl = string(oauthTokenUrlData)
	}
	enrichedContext = context.WithValue(enrichedContext, oauthCTokenUrlContextKey, oauthTokenUrl)

	// TODO add missing scopes
	//scopes := strings.Split(oauthScopesString.(string), ",")
	//var massagedScopes []string
	//for _, scope := range scopes {
	//	massagedScopes = append(massagedScopes, strings.Trim(scope, " "))
	//}

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
