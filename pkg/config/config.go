package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type config struct {
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
}
type server struct {
	Httpport int `mapstructure:"httpport"`
}
type database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBname   string `mapstructure:"dbname"`
}

var Config config

func init() {
	v := viper.New()
	v.SetConfigType("ini")
	v.AddConfigPath("conf")
	v.SetConfigName("app.ini")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	settings := v.AllSettings()
	err := mapstructure.WeakDecode(settings, &Config)
	if err != nil {
		panic(err)
	}
}
