package auth

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/BaytoorJr/shared-libs/errors"
)

// Config struct & variable
var authConfig *Config

type Config struct {
	ProjectSlug string `envconfig:"project_slug" required:"true"`
	AuthConnURL string `envconfig:"auth_conn_url" required:"true"`
	RoleConnURL string `envconfig:"role_conn_url" required:"true"`
}

// Configs getter
func getAuthConfig() (*Config, error) {
	if authConfig != nil {
		return authConfig, nil
	}

	config, err := readEnvConfigs()
	if err != nil {
		return nil, err
	}

	authConfig = config

	return authConfig, nil
}

// Read configs from ENV
func readEnvConfigs() (*Config, error) {
	var config Config

	err := envconfig.Process("SSO_CONFIG", &config)
	if err != nil {
		return nil, errors.ENVReadError.SetDevMessage(err.Error())
	}

	return &config, nil
}
