package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	SupaBaseConnection string `mapstructure:"Supabaseconnection"`
	JwtSecret          string `mapstructure:"jwtsecret"`
	ConfigFilePath     string `mapstructure:"configFilePath"`
	ConfigPath         string `mapstructure:"configPath"`
	EncrytpSecret      string `mapstructure:"secret"`
}

func InitConfig() (*Config, error) {
	var config = Config{}

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return &config, err
}
