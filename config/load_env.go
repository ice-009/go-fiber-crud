package config

import("github.com/spf13/viper")

type Config struct {
	DBHost string `mapstructure:"POSGRES_HOST"`
	DBUsername string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`
	DBPort string `mapstructure:"POSTGRES_PORT"`
}

func LoadConfig(path string)(config Config, err error){
	viper.addConfigPath(path);
	viper.SetConfigName("app");
	viper.SetConfigType("env");

	viper.AutomaticEnv();

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}