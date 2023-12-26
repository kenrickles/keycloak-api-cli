package cmd

import (
	"fmt"
	"keycloak-api-cli/pkg/keycloak"
	
	"github.com/spf13/cobra"
)


// GetCommand creates and returns a new list command
func RetrieveCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
	RetrieveCommand := &cobra.Command{
		Use:   "retrieve [token]",
		Short: "retrieve token in Keycloak",
	}
	RetrieveCommand.AddCommand(RetrieveTokenCommand(kcClient))

	return RetrieveCommand
}

// Retrieve Token Command
func RetrieveTokenCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    return &cobra.Command{
        Use:   "token",
        Short: "Retrieve token in Keycloak",
        Run: func(cmd *cobra.Command, args []string) {
            if kcClient.Token == "" {
                fmt.Println("Error: Token doesn't exist")
                return
            }
            fmt.Printf(kcClient.Token)
        },
    }
}