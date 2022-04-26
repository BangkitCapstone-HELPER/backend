package config

import (
	"time"
)

type jwtConfig struct {
	JWTConfig *jwtConfigImpl `mapstructure:"JWT"`
}

type jwtConfigImpl struct {
	Algorithm        string `mapstructure:"Algorithm"`
	SignatureKey     string `mapstructure:"Key"`
	ExpirationInHour int    `mapstructure:"Exp"`
}

type JWTConfig interface {
	Exp() time.Duration
	Secret() string
}

func NewJWT(viperLoader ViperLoader) JWTConfig {
	cfg := &jwtConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *jwtConfig) Exp() time.Duration {
	return time.Duration(a.JWTConfig.ExpirationInHour) * time.Hour
}

func (a *jwtConfig) Secret() string {
	return a.JWTConfig.SignatureKey
}
