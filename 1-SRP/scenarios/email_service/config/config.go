package config

import "github.com/spf13/viper"

func Init() error {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")        // path to look for the config file in
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	return viper.ReadInConfig()
}
