package config

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// viperLoader is a struct that implements Loader interface
type viperLoader struct{}

// ViperLoader is a function that returns a new instance of viperLoader
type ViperLoader interface {
	Unmarshal(i interface{})
}

var viperLoadOnce sync.Once

func NewViperLoader() ViperLoader {
	viperLoadOnce.Do(func() {
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			panic(errors.Wrap(err, "failed to read config"))
		}
	})

	return &viperLoader{}
}

func (v viperLoader) Unmarshal(i interface{}) {
	if err := viper.Unmarshal(&i); err != nil {
		panic(errors.Wrap(err, "failed to marshal config"))
	}
}
