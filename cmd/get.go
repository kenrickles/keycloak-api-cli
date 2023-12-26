package cmd

import (
	"fmt"
	"keycloak-api-cli/pkg/keycloak"
	
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
)


// GetCommand creates and returns a new list command
func GetCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
	GetCommand := &cobra.Command{
		Use:   "get [users, realms, resources, roles]",
		Short: "List users, realms, resources, or roles in Keycloak",
	}
	GetCommand.AddCommand(GetRealmCommand(kcClient))
	GetCommand.AddCommand(GetUsersCommand(kcClient))

	return GetCommand
}

// List Realm
func GetRealmCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
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

func GetUsersCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    var realm string
    usersCmd := &cobra.Command{
        Use:   "users",
        Short: "List all users",
        Long:  `List all users in the Keycloak instance.`,
        Run: func(cmd *cobra.Command, args []string) {
            // Check if the realm flag is set, else use the default realm
            realm, _ := cmd.Flags().GetString("realm")
            if realm == "" {
                realm = viper.GetString("default_realm")
            }

            // Use the specified or default realm for listing users
            users, err := kcClient.ListUsers(realm) // Make sure ListUsers accepts a realm parameter
            if err != nil {
                fmt.Printf("Error listing users in realm %s: %v\n", realm, err)
                return
            }

            // Iterate and print details of each user
            for _, user := range users {
                fmt.Printf("Realm: %s, User Name: %s, User ID: %s\n", realm, user.Username, user.ID)
            }
        },
    }

    // Attach the realm flag to the command
    usersCmd.Flags().StringVarP(&realm, "realm", "r", "", "Specify the realm to list users")

    return usersCmd
}