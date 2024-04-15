package config

// ServerConfig 项目配置
type ServerConfig struct {
	ServiceInfo Service     `mapstructure:"service" json:"service" yaml:"service"`
	MysqlInfo   MySqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
	LogInfo     LogConfig   `mapstructure:"log" json:"log" yaml:"log"`
}

// Service 外部服务地址配置
type Service struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

// MySqlConfig 数据源配置
type MySqlConfig struct {
	UserName     string `mapstructure:"username" json:"username" yaml:"username"` // mysql 数据库账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // mysql数据库密码
	Host         string `mapstructure:"host" json:"host" yaml:"host"`             // mysql主机地址
	Port         int64  `mapstructure:"port" json:"port" yaml:"port"`             // 端口号
	DbName       string `mapstructure:"dbname" json:"dbname" yaml:"dbName"`       // 连接的数据库名称
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

// RedisConfig Redis配置
type RedisConfig struct {
	Database int    `mapstructure:"database" json:"database" yaml:"database"` // 选择redis几号库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 端口号
	Password string `mapstructure:"password" json:"password" yaml:"password"` // redis连接密码
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `json:"level" yaml:"level"`
	Path     string `json:"path" yaml:"path"`
	Filename string `json:"filename" yaml:"filename"`
}
