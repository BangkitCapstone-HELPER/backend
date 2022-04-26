package config

import "fmt"

type httpConfig struct {
	HTTPConfig *httpConfigImpl `mapstructure:"HTTP"`
}

type httpConfigImpl struct {
	Host string `mapstructure:"Host" validate:"ipv4"`
	Port int    `mapstructure:"Port" validate:"gte=1,lte=65535"`
}

// HTTPConfig is the configuration for the HTTP server
type HTTPConfig interface {
	ListenAddr() string
}

// NewHTTP creates a new HTTPConfig
func NewHTTP(viperLoader ViperLoader) HTTPConfig {
	cfg := &httpConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *httpConfig) ListenAddr() string {
	return fmt.Sprintf("%s:%d", a.HTTPConfig.Host, a.HTTPConfig.Port)
}
