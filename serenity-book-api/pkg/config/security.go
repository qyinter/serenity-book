package config

// Security 定义了安全性相关配置
type Security struct {
	APISign     APISignConfig `mapstructure:"api_sign" json:"api_sign" yaml:"api_sign"`
	Secret      string        `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expire      int           `mapstructure:"expire" json:"expire" yaml:"expire"`
	TokenPrefix string        `mapstructure:"token_prefix" json:"token_prefix" yaml:"token_prefix"`
}

// APISignConfig 定义了 API 签名的相关配置
type APISignConfig struct {
	AppKey      string `mapstructure:"app_key" json:"app_key" yaml:"app_key"`
	AppSecurity string `mapstructure:"app_security" json:"app_security" yaml:"app_security"`
}
