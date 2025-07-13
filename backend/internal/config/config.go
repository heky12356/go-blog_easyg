package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		JWTKEY string `mapstructure:"jwtkey"`
	} `mapstructure:"app"`
	DateBase struct {
		Litedb string `mapstructure:"litedb"`
	} `mapstructure:"database"`
}

var config = &Config{}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}

func GetConfig() *Config { return config }
