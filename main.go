package main

import (
    "log"
    "keycloak-api-cli/cmd"
    "keycloak-api-cli/pkg/keycloak"
    "keycloak-api-cli/pkg/config"
)

func main() {
    // Load your configuration (modify as per your configuration logic)
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize Keycloak client
    kcClient := keycloak.NewClient(cfg)

    // Execute the root command with the Keycloak client
    cmd.Execute(kcClient)
}
