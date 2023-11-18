package config

import (
    "github.com/spf13/viper"
)

type KeycloakConfig struct {
    URL          string `mapstructure:"url"`
    ClientID     string `mapstructure:"client_id"`
    ClientSecret string `mapstructure:"client_secret"`
    Username     string `mapstructure:"username"`
    Password     string `mapstructure:"password"`
    AuthRealm    string `mapstructure:"auth_realm"`
    DefaultRealm  string `mapstructure:"default_realm"`
}

// LoadConfig loads configuration from config.yaml
func LoadConfig() (KeycloakConfig, error) {
    var config KeycloakConfig

    // Set default configurations 	
	viper.SetConfigName("config") 
    viper.AddConfigPath(".")       
    viper.AutomaticEnv()
    if err := viper.ReadInConfig(); err != nil {
        return config, err
    }

    err := viper.Unmarshal(&config)
    return config, err
}
