package util

import "github.com/spf13/viper"

type GlobalConfig struct {
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBSource               string `mapstructure:"DB_SOURCE"`
	FdAuthServerAddress    string `mapstructure:"FD_AUTH_SERVER_ADDRESS"`
	FdEventListenerAddress string `mapstructure:"FD_EVENT_LISTENER_SERVER_ADDRESS"`
}

func LoadGlobalConfig(path string) (config GlobalConfig, err error) {
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
