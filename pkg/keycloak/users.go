package keycloak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User represents a Keycloak user
type User struct {
	// Define user properties
}

// ListUsers lists all users in Keycloak
func ListUsers() {
	url := "http://<your-keycloak-url>/admin/realms/<your-realm>/users"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer <your-access-token>")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on response.\n[ERROR] -", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var users []User
	json.Unmarshal(body, &users)

	// Print users or handle them as needed
	fmt.Println("Users:", users)
}
