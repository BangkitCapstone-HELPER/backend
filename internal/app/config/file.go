package config

type fileConfig struct {
	FileConfig *fileConfigImpl `mapstructure:"File"`
}

type fileConfigImpl struct {
	Path string `mapstructure:"StoragePath"`
}

type FileConfig interface {
	Path() string
}

func NewFile(viperLoader ViperLoader) FileConfig {
	cfg := &fileConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *fileConfig) Path() string {
	return a.FileConfig.Path
}
