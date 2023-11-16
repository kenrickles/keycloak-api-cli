package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// Manage User (keycloak-api-cli <ACTION>)
var usersCmd = &cobra.Command{
    Use:   "users",
    Short: "Manage users in Keycloak",
    Long:  `This command allows you to manage users in Keycloak.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Users command called")
    },
}

func init() {
    rootCmd.AddCommand(usersCmd)
}
