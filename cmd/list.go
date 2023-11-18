package cmd

import (
	"fmt"
	"keycloak-api-cli/pkg/keycloak"
	
	"github.com/spf13/cobra"
)


// ListCommand creates and returns a new list command
func ListCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
	ListCommand := &cobra.Command{
		Use:   "list [users, realms, resources, roles]",
		Short: "List users, realms, resources, or roles in Keycloak",
	}
	ListCommand.AddCommand(ListRealmCommand(kcClient))

	return ListCommand
}



// List Realm
func ListRealmCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    return &cobra.Command{
        Use:   "realms",
        Short: "List all realms",
        Long:  `List all realms in the Keycloak instance.`,
        Run: func(cmd *cobra.Command, args []string) {
            realms, err := kcClient.ListRealms()
            if err != nil {
                fmt.Printf("Error listing realms: %v\n", err)
                return
            }
            for _, realm := range realms {
				fmt.Printf("Realm Name: %s, Realm ID: %s\n", realm.Realm, realm.ID)
            }
        },
    }
}