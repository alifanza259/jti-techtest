package util

import (
	"github.com/spf13/viper"
)

// In order to get the value of the variables and store them in this struct, we need to use the unmarshaling feature of Viper.
// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBName string `mapstructure:"DB_NAME"`
	DBPass string `mapstructure:"DB_PASS"`
	Host   string `mapstructure:"HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
