package config

import "fmt"

type databaseConfig struct {
	DatabaseConfig *databaseConfigImpl `mapstructure:"Database"`
}

type databaseConfigImpl struct {
	Engine      string `mapstructure:"Engine"`
	Name        string `mapstructure:"Name"`
	Host        string `mapstructure:"Host"`
	Port        int    `mapstructure:"Port"`
	Username    string `mapstructure:"Username"`
	Password    string `mapstructure:"Password"`
	TablePrefix string `mapstructure:"TablePrefix"`
	Parameters  string `mapstructure:"Parameters"`

	MaxLifetime  int `mapstructure:"MaxLifetime"`
	MaxOpenConns int `mapstructure:"MaxOpenConns"`
	MaxIdleConns int `mapstructure:"MaxIdleConns"`
}

type DatabaseConfig interface {
	Engine() string
	DSN() string
}

func NewDatabase(viperLoader ViperLoader) DatabaseConfig {
	cfg := &databaseConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *databaseConfig) Engine() string {
	return a.DatabaseConfig.Engine
}

func (a *databaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", a.DatabaseConfig.Host, a.DatabaseConfig.Username, a.DatabaseConfig.Password, a.DatabaseConfig.Name, a.DatabaseConfig.Port)
}
