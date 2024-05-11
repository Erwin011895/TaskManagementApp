package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/ini"
)

type Config struct {
	App      App      `mapstructure:"App"`
	Postgres Postgres `mapstructure:"Postgres"`
}

type App struct {
	Key  string `mapstructure:"Key"`
	Host string `mapstructure:"Host"`
	Port string `mapstructure:"Port"`
}

type Postgres struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Database string `mapstructure:"Database"`
}

func SetupConfig() (conf Config) {
	config.AddDriver(ini.Driver)

	err := config.LoadFiles("config.ini")
	if err != nil {
		panic(err)
	}

	err = config.Decode(&conf)
	if err != nil {
		panic(err)
	}

	return
}
