package config

type Server struct {
	// 验证码配置
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
