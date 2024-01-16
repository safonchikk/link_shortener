package util

import "github.com/spf13/viper"

type Config struct {
	RedisAddr   string `mapstructure:"REDIS_ADDR"`
	AppPort     string `mapstructure:"APP_PORT"`
	NginxPort   string `mapstructure:"NGINX_PORT"`
	Address     string `mapstructure:"ADDRESS"`
	LinkExpTime string `mapstructure:"LINK_EXP_TIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
