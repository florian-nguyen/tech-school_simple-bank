package util

import "github.com/spf13/viper"

// Config stores all the configuration of the application. The values are read by viper from a config file of environment variable.
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads the configuration from file or the environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // because name of config file is app.env
	viper.SetConfigType("env") // because file extension is .env

	viper.AutomaticEnv() // will override all default environment variables with the ones found in the input file

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
