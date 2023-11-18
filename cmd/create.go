package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"keycloak-api-cli/pkg/keycloak"

	"github.com/spf13/cobra"
)



// CreateCommand creates and returns a new create command
func CreateCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
	CreateCommand := &cobra.Command{
		Use:   "create [user, realm, resource, role]",
		Short: "Create a user, realm, resource, or role in Keycloak",
	}
	CreateCommand.AddCommand(CreateRealmCommand(kcClient))
	
	return CreateCommand
}

// Create Realm
func CreateRealmCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    CreateRealmCommand := &cobra.Command{
        Use:   "realm",
        Short: "Create a new Keycloak realm",
        Run: func(cmd *cobra.Command, args []string) {
            if realmName == "" {
                realmName = askForRealmName()
            }
            if err := kcClient.CreateRealm(realmName); err != nil {
                fmt.Printf("Error creating realm: %v\n", err)
            } else {
                fmt.Println("Realm", realmName, "created successfully")
            }
        },
    }
    CreateRealmCommand.Flags().StringVarP(&realmName, "name", "n", "", "Name of the realm to create")
    return CreateRealmCommand
}

func askForRealmName() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the name of the realm to create: ")
    name, _ := reader.ReadString('\n')
    return strings.TrimSpace(name)
}