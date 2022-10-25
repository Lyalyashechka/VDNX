package config

import "github.com/spf13/viper"

type Config struct {
	Postgres Postgres `mapstructure:"postgres"`
}

type Postgres struct {
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

func LoadConfig(config *Config, path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	return nil
}
