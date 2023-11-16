package cmd

import (
    "fmt"
	"bufio"
	"os"
	"strings"
    "keycloak-api-cli/pkg/keycloak"

    "github.com/spf13/cobra"
)
// Intialise Variable
var realmName string 

// Create new Realm
func realmsCmd(kcClient *keycloak.KeycloakClient) *cobra.Command {
    realmsCmd := &cobra.Command{
        Use:   "realms",
        Short: " List, Create, Delete Keycloak realms",
    }

    // Sub commands
    realmsCmd.AddCommand(listCmd(kcClient))
    realmsCmd.AddCommand(createCmd(kcClient))

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
				fmt.Printf("Realm Name: %s, Realm ID: %s\n", realm.Realm, realm.ID)
            }
        },
    }
}



// Create Realm
func createCmd(kcClient *keycloak.KeycloakClient) *cobra.Command {
    createRealmCmd := &cobra.Command{
        Use:   "create",
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
    createRealmCmd.Flags().StringVarP(&realmName, "name", "n", "", "Name of the realm to create")
    return createRealmCmd
}

func askForRealmName() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the name of the realm to create: ")
    name, _ := reader.ReadString('\n')
    return strings.TrimSpace(name)
}
