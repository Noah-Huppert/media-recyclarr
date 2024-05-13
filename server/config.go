package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config is application config
type Config struct {
	// EmbyURL is the URL to the Emby server
	EmbyURL string `env:"EMBY_URL,required"`

	// EmbyAPIKey is the API key used to authenticate with Emby
	EmbyAPIKey string `env:"EMBY_API_KEY,required"`

	// JellyseerrURL is the URL to the Jellyseerr server
	JellyseerrURL string `env:"JELLYSEERR_URL,required"`

	// JellyseerrAPIKey is the API key used to authenticate with Jellyseerr
	JellyseerrAPIKey string `env:"JELLYSEERR_API_KEY,required"`
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.ParseWithOptions(&cfg, env.Options{
		Prefix: "MEDIA_RECYCLARR_",
	}); err != nil {
		return nil, fmt.Errorf("failed to load from environment variables: %s", err)
	}

	return &cfg, nil
}
