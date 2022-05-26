package config

type bucketConfig struct {
	BucketConfig *bucketConfigImpl `mapstructure:"CloudStorage"`
}

type bucketConfigImpl struct {
	BucketName string `mapstructure:"BucketName"`
	ProjectID  string `mapstructure:"ProjectID"`
	Path       string `mapstructure:"Path"`
}

type BucketConfig interface {
	BucketName() string
	ProjectID() string
	Path() string
}

func NewBucket(viperLoader ViperLoader) BucketConfig {
	cfg := &bucketConfig{}

	viperLoader.Unmarshal(cfg)

	return cfg
}

func (a *bucketConfig) BucketName() string {
	return a.BucketConfig.BucketName
}

func (a *bucketConfig) ProjectID() string {
	return a.BucketConfig.ProjectID
}

func (a *bucketConfig) Path() string {
	return a.BucketConfig.Path
}
