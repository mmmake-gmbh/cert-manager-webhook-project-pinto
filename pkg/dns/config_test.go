package dns

import (
	"context"
	"gitlab.com/whizus/customer/pinto/cert-manager-webhook-pinto/internal/testutils"
	v1 "k8s.io/api/core/v1"
	extapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"reflect"
	"testing"
)

func TestConfig_Environment(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultEnvironment,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), environmentContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.Environment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Environment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Name(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{{
		name: "default value is returned",
		fields: fields{
			savedContext: context.Background(),
		},
		want: "pinto",
	},
	// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_OauthClientID(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultOAuthClientID,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), oauthClientIDContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.OauthClientID(); got != tt.want {
				t.Errorf("OauthClientID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_OauthClientScopes(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}
	testScopes := []string{
		"scope1",
		"scope2",
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultOauthScopes,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), oauthScopesContextKey, testScopes),
			},
			want: testScopes,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.OauthClientScopes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OauthClientScopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_OauthClientSecret(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultOAuthClientSecret,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), oauthClientSecretContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.OauthClientSecret(); got != tt.want {
				t.Errorf("OauthClientSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_OauthTokenURL(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultOAuthTokenURL,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), oauthCTokenUrlContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.OauthTokenURL(); got != tt.want {
				t.Errorf("OauthTokenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_PintoApiURL(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultPintoApiURL,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), pintoApiUrlContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.PintoApiURL(); got != tt.want {
				t.Errorf("PintoApiURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Provider(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testValueString := "test_value"

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default value is returned",
			fields: fields{
				savedContext: context.Background(),
			},
			want: defaultProvider,
		},
		{
			name: "set string value is returned",
			fields: fields{
				savedContext: context.WithValue(context.Background(), providerContextKey, testValueString),
			},
			want: testValueString,
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.Provider(); got != tt.want {
				t.Errorf("Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_getContext(t *testing.T) {
	type fields struct {
		savedContext context.Context
	}

	testContext := context.Background()

	tests := []struct {
		name   string
		fields fields
		want   context.Context
	}{
		{
			name:   "returns nil if not defined",
			fields: fields{},
			want:   nil,
		},
		{
			name: "returns same context if defined",
			fields: fields{
				savedContext: testContext,
			},
			want: testContext,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				savedContext: tt.fields.savedContext,
			}
			if got := c.getContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		cfgJSON *extapi.JSON
	}
	// prepare fixtures
	configBytes, readErr := testutils.ReadFixture("config.json")
	if readErr != nil {
		t.Error(readErr)
		return
	}
	testJson := new(extapi.JSON)
	testJson.Raw = configBytes

	testConfig := ProviderConfig{}
	testConfig.AccessKey = new(v1.SecretKeySelector)
	testConfig.AccessKey.Name = "test-secret"
	testConfig.AccessKey.Key = "PINTO_OAUTH_CLIENT_ID"
	testConfig.SecretKey = new(v1.SecretKeySelector)
	testConfig.SecretKey.Name = "test-secret"
	testConfig.SecretKey.Key = "PINTO_OAUTH_CLIENT_SECRET"
	testConfig.PintoApiUrl = new(v1.SecretKeySelector)
	testConfig.PintoApiUrl.Name = "test-secret"
	testConfig.PintoApiUrl.Key = "PINTO_API_URL"
	testConfig.OauthTokenUrl = new(v1.SecretKeySelector)
	testConfig.OauthTokenUrl.Name = "test-secret"
	testConfig.OauthTokenUrl.Key = "PINTO_OAUTH_TOKEN_URL"

	tests := []struct {
		name    string
		args    args
		want    ProviderConfig
		wantErr bool
	}{
		{
			name: "return empty config if parameter is nil",
			args: args{
				cfgJSON: nil,
			},
			want:    ProviderConfig{},
			wantErr: false,
		},
		{
			name: "load full config file",
			args: args{
				cfgJSON: testJson,
			},
			want:    testConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.cfgJSON)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
