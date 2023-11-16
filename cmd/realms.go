package cmd

import (
    "fmt"
    "keycloak-api-cli/pkg/keycloak"

    "github.com/spf13/cobra"
)
// Create new Realm
func NewRealmsCmd(kcClient *keycloak.KeycloakClient) *cobra.Command {
    realmsCmd := &cobra.Command{
        Use:   "realms",
        Short: "Manage Keycloak realms",
    }

    // Sub commands
    realmsCmd.AddCommand(listCmd(kcClient))

    return realmsCmd
}

// List Realm
func listCmd(kcClient *keycloak.KeycloakClient) *cobra.Command {
    return &cobra.Command{
        Use:   "list",
        Short: "List all realms",
        Long:  `List all realms in the Keycloak instance.`,
        Run: func(cmd *cobra.Command, args []string) {
            realms, err := kcClient.ListRealms()
            if err != nil {
                fmt.Printf("Error listing realms: %v\n", err)
                return
            }
            for _, realm := range realms {
                fmt.Println(realm.ID, realm.DisplayName)
            }
        },
    }
}
