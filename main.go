package main

import (
    "flag"
    "log"
    "keycloak-api-cli/cmd"
    "keycloak-api-cli/pkg/keycloak"
    "keycloak-api-cli/pkg/config"
)

func main() {
    // Define a command-line flag for the config file
    var cfgFile string
    flag.StringVar(&cfgFile, "config", "config.yaml", "path to config file")
    flag.Parse()

    // Load your configuration using the specified config file
    cfg, err := config.LoadConfig(cfgFile)
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize Keycloak client
    kcClient := keycloak.NewClient(cfg)
    err = kcClient.Authenticate(cfg)
    if err != nil {
        log.Fatalf("Failed to authenticate: %v", err)
    }

    // Execute the root command with the Keycloak client
    cmd.Execute(kcClient)
}
