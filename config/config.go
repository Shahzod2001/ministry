package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Redis    Redis    `mapstructure:"redis"`
	TokenTTL TokenTTL `mapstructure:"token_ttl"`
	Logger   Logger   `mapstructure:"logger"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type TokenTTL struct {
	AccessTokenTTL  string `mapstructure:"access_token_ttl"`
	RefreshTokenTTL string `mapstructure:"refresh_token_ttl"`
}

type Logger struct {
	Format         string `mapstructure:"format"`
	WriteToFile    bool   `mapstructure:"write_to_file"`
	WriteToConsole bool   `mapstructure:"write_to_console"`
	LogFile        string `mapstructure:"log_file"`
}

const (
	configName = "config"
	configType = "yml"
	configPath = "config"
)

var AppParams *Config

func Init() *Config {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	AppParams = &cfg

	return &cfg
}
