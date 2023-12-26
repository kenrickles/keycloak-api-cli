package config

import (
    "fmt"
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
func LoadConfig(configFile string) (KeycloakConfig, error) {
    var config KeycloakConfig

    viper.SetConfigFile(configFile) // Set the configuration file
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return config, fmt.Errorf("failed to load config file '%s': %v", configFile, err)
    }

    err := viper.Unmarshal(&config)
    if err != nil {
        return config, fmt.Errorf("error unmarshalling config: %v", err)
    }

    return config, nil
}