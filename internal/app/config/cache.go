package config

import "fmt"

type cacheConfig struct {
	CacheConfig *cacheConfigImpl `mapstructure:"Cache"`
}

type cacheConfigImpl struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Password string `mapstructure:"Password"`
	Database string `mapstructure:"Database"`
}

type CacheConfig interface {
	Address() string
	Password() string
	Database() string
}

func NewCache(viperLoader ViperLoader) CacheConfig {
	cfg := &cacheConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *cacheConfig) Password() string {
	return a.CacheConfig.Password
}

func (a *cacheConfig) Database() string {
	return a.CacheConfig.Database
}

func (a *cacheConfig) Address() string {
	return fmt.Sprintf("%s:%d", a.CacheConfig.Host, a.CacheConfig.Port)
}
