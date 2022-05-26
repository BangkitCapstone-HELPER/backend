package config

type emailConfig struct {
	emailConfig *emailConfigImpl `mapstructure:"Email"`
}

type emailConfigImpl struct {
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Sender   string `mapstructure:"Sender"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
}

type EmailConfig interface {
	Host() string
	Port() int
	Sender() string
	Password() string
	Username() string
}

func NewEmail(viperLoader ViperLoader) EmailConfig {
	cfg := &emailConfig{}

	viperLoader.Unmarshal(cfg)
	return cfg
}

func (a *emailConfig) Host() string {
	return a.emailConfig.Host
}
func (a *emailConfig) Port() int {
	return a.emailConfig.Port
}
func (a *emailConfig) Sender() string {
	return "Rahmat Wibowo <rahmat.wibowo21@gmail.com>"
}
func (a *emailConfig) Password() string {
	return a.emailConfig.Password
}
func (a *emailConfig) Username() string {
	return a.emailConfig.Username
}
