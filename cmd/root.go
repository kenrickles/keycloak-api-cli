package cmd

import (
    "fmt"
    "os"
    "keycloak-api-cli/pkg/keycloak"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
    Use:   "keycloak-api-cli",
    Short: "CLI to interact with Keycloak",
    Long:  `A Command Line Interface to interact with Keycloak API`,
    // ...
}

func Execute(kcClient *keycloak.KeycloakClient) {
    rootCmd.AddCommand(NewRealmsCmd(kcClient))

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
	cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.yaml", "config file (default is $HOME/.keycloak-api-cli.yaml)")

}

func initConfig() {
    fmt.Println("Initializing configuration...")

    if cfgFile != "" {
        fmt.Println("Using config file:", cfgFile)
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := os.UserHomeDir()
        if err != nil {
            fmt.Println("Unable to find home directory:", err)
            os.Exit(1)
        }
        fmt.Println("Looking for config file in home directory:", home)
        viper.AddConfigPath(home)
        viper.SetConfigName("config.yaml")
    }

    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("Failed to read config file:", err)
        os.Exit(1)
    } else {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
