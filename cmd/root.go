package cmd

import (
    "fmt"
    "os"
    "keycloak-api-cli/pkg/keycloak"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// Intialise Variable
var cfgFile string
var realmName string
var username string
var password string
var email string

var rootCmd = &cobra.Command{
    Use:   "keycloak-api-cli",
    Short: "CLI to interact with Keycloak",
    Long:  `A Command Line Interface to interact with Keycloak API`,
    // ...
}

func Execute(kcClient *keycloak.KeycloakClient) {
    rootCmd.AddCommand(RetrieveCommand(kcClient))
    rootCmd.AddCommand(GetCommand(kcClient))
    rootCmd.AddCommand(CreateCommand(kcClient))
    rootCmd.AddCommand(DeleteCommand(kcClient))

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    // Set the default to an empty string
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initConfig() {
    if cfgFile != "" {
        // Use the specified config file
        fmt.Println("Using specified config file:", cfgFile)
        viper.SetConfigFile(cfgFile)
    } else {
        defaultConfig := "config.yaml"
        fmt.Println("Using default config file:", defaultConfig)
        viper.SetConfigFile(defaultConfig)
    }

    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("Failed to read config file:", err)
        os.Exit(1)
    } else {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
