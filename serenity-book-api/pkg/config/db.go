package config

import (
	"time"
)

// DatabaseConfig holds the MySQL database configuration.
type DatabaseConfig struct {
	Username  string `yaml:"username" mapstructure:"username"`
	Password  string `yaml:"password" mapstructure:"password"`
	Host      string `yaml:"host" mapstructure:"host"`
	Port      string `yaml:"port" mapstructure:"port"`
	Database  string `yaml:"database" mapstructure:"database"`
	Charset   string `yaml:"charset" mapstructure:"charset"`
	ParseTime bool   `yaml:"parseTime" mapstructure:"parseTime"`
	Loc       string `yaml:"loc" mapstructure:"loc"`
}

// RedisConfig holds the Redis connection configuration.
type RedisConfig struct {
	Addr         string        `yaml:"addr" mapstructure:"addr"`
	Password     string        `yaml:"password" mapstructure:"password"`
	DB           int           `yaml:"db" mapstructure:"db"`
	ReadTimeout  time.Duration `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout" mapstructure:"write_timeout"`
}

// DbConfig holds the entire configuration.
type DbConfig struct {
	Mysql DatabaseConfig `yaml:"mysql" mapstructure:"mysql"`
	Redis RedisConfig    `yaml:"redis" mapstructure:"redis"`
}
