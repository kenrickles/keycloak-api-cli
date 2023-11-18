package cmd

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"keycloak-api-cli/pkg/keycloak"

	"github.com/spf13/cobra"
)


// DeleteCommand creates and returns a new delete command
func DeleteCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    DeleteCommand := &cobra.Command{
        Use:   "delete [user, realm, resource, role]",
        Short: "Delete a Keycloak realm",
    }
    DeleteCommand.AddCommand(DeleteRealmCommand(kcClient))
    DeleteCommand.AddCommand(DeleteUserCommand(kcClient))
    
    return DeleteCommand
}

// Delete Realm
func DeleteRealmCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    DeleteRealmCommand := &cobra.Command{
        Use:   "realm",
        Short: "Delete a Keycloak realm by name",
        Run: func(cmd *cobra.Command, args []string) {
            if realmName == "" {
                realmName = askForRealmName()
            }
            // You can add confirmation logic here to confirm deletion
            fmt.Printf("Are you sure you want to delete the realm %s? (yes/no): ", realmName)
            reader := bufio.NewReader(os.Stdin)
            confirmation, _ := reader.ReadString('\n')
            confirmation = strings.TrimSpace(confirmation)
            
            if confirmation != "yes" {
                fmt.Println("Deletion canceled.")
                return
            }
            
            if err := kcClient.DeleteRealm(realmName); err != nil {
                fmt.Printf("Error deleting realm: %v\n", err)
            } else {
                fmt.Printf("Realm %s deleted successfully\n", realmName)
            }
        },
    }
    DeleteRealmCommand.Flags().StringVarP(&realmName, "name", "n", "", "Name of the realm to delete")
    return DeleteRealmCommand
}

// Delete User
func DeleteUserCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    var (
        userID   string
        username string
    )

    DeleteUserCommand := &cobra.Command{
        Use:   "user",
        Short: "Delete a Keycloak user by username",
        Run: func(cmd *cobra.Command, args []string) {
            if username == "" && userID == "" {
                username = askForUserName()
            }

            if username != "" {
                // Use username to retrieve userID
                retrievedUserID, err := kcClient.GetUserIDByUsername(username)
                if err != nil {
                    fmt.Printf("Error retrieving userID for username %s: %v\n", username, err)
                    return
                }
                userID = retrievedUserID
            }

            // You can add confirmation logic here to confirm deletion
            fmt.Printf("Are you sure you want to delete the user %s (userID: %s)? (yes/no): ", username, userID)
            reader := bufio.NewReader(os.Stdin)
            confirmation, _ := reader.ReadString('\n')
            confirmation = strings.TrimSpace(confirmation)

            if confirmation != "yes" {
                fmt.Println("Deletion canceled.")
                return
            }

            if err := kcClient.DeleteUser(userID); err != nil {
                fmt.Printf("Error deleting user: %v\n", err)
            } else {
                fmt.Printf("User %s (userID: %s) deleted successfully\n", username, userID)
            }
        },
    }
    DeleteUserCommand.Flags().StringVarP(&userID, "userid", "i", "", "ID of the user to delete")
    DeleteUserCommand.Flags().StringVarP(&username, "username", "u", "", "Username of the user to delete")
    return DeleteUserCommand
}
