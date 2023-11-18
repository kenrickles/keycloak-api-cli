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
	CreateCommand.AddCommand(CreateUserCommand(kcClient))
	
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


// Create user
func CreateUserCommand(kcClient *keycloak.KeycloakClient) *cobra.Command {
    CreateUserCommand := &cobra.Command{
        Use:   "user",
        Short: "Create a new Keycloak user",
        Run: func(cmd *cobra.Command, args []string) {
            if username == "" {
                username = askForUserName()
            }
			if password == "" {
                password = askForUserPassword()
            }
			if email == "" {
                email = askForUserEmail()
            }
            if err := kcClient.CreateUser(username, password, email); err != nil {
                fmt.Printf("Error creating user: %v\n", err)
            } else {
                fmt.Println("username", username, "created successfully")
            }
        },
    }
    CreateUserCommand.Flags().StringVarP(&realmName, "username", "u", "", "username of the user to create")
	CreateUserCommand.Flags().StringVarP(&password, "password", "p", "", "Password for the user")
	CreateUserCommand.Flags().StringVarP(&email, "email", "e", "", "Email address of the user")
    return CreateUserCommand
}

func askForUserName() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the name of the user: ")
    name, _ := reader.ReadString('\n')
    return strings.TrimSpace(name)
}

func askForUserPassword() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the password for the user: ")
    password, _ := reader.ReadString('\n')
    return strings.TrimSpace(password)
}

func askForUserEmail() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the email address of the user: ")
    email, _ := reader.ReadString('\n')
    return strings.TrimSpace(email)
}

func askForUserID() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the ID of the user: ")
    email, _ := reader.ReadString('\n')
    return strings.TrimSpace(email)
}