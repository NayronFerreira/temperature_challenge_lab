package config

import "github.com/spf13/viper"

type Config struct {
	WebServerPort  string `mapstructure:"WEB_SERVER_PORT"`
	ViaCepHost     string `mapstructure:"VIACEP_HOST_API"`
	WeatherHost    string `mapstructure:"WEATHER_HOST_API"`
	WeatherKey     string `mapstructure:"WEATHER_API_KEY"`
	TimeoutSeconds string `mapstructure:"TIMEOUT_SECONDS"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile("secrets.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config, err
}
