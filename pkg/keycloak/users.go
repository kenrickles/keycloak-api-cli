package keycloak

import (
    "encoding/json"
    "bytes"
    "fmt"
	"io/ioutil"
    "net/http"
)

// User represents a Keycloak user
type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`

}

// CreateUser creates a new user in Keycloak
func (kc *KeycloakClient) CreateUser(realm, username, email, password string) error {
    userDetails := struct {
        Username    string `json:"username"`
        Email       string `json:"email"`
        Credentials []struct {
            Type      string `json:"type"`
            Value     string `json:"value"`
            Temporary bool   `json:"temporary"`
        } `json:"credentials"`
    }{
        Username: username,
        Email:    email,
        Credentials: []struct {
            Type      string `json:"type"`
            Value     string `json:"value"`
            Temporary bool   `json:"temporary"`
        }{
            {
                Type:      "password",
                Value:     password,
                Temporary: false, // Set to true if you want the user to change password on first login
            },
        },
    }

    // Use the provided realm in the URL instead of the default realm
    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, realm)

    // Encode the user details as JSON for the request body
    jsonBody, err := json.Marshal(userDetails)
    if err != nil {
        return err
    }

    // Create a new POST request with the JSON body
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+kc.Token)
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    res, err := kc.Client.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    // Check for successful status code, typically 201 for creation
    if res.StatusCode != http.StatusCreated {
        errorResponse, _ := ioutil.ReadAll(res.Body)
        return fmt.Errorf("failed to create user in realm %s, received status code: %d, error message: %s", realm, res.StatusCode, string(errorResponse))
    }

    return nil
}


func (kc *KeycloakClient) ListUsers(realm string) ([]User, error) {
    // Use the provided realm in the URL instead of the default realm
    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, realm)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+kc.Token)

    res, err := kc.Client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var users []User
    err = json.NewDecoder(res.Body).Decode(&users)
    if err != nil {
        return nil, err
    }

    return users, nil
}


func (kc *KeycloakClient) DeleteUser(realm, userID string) error {
    url := fmt.Sprintf("%s/admin/realms/%s/users/%s", kc.BaseURL, realm, userID)

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+kc.Token)

    res, err := kc.Client.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusNoContent {
        errorResponse, _ := ioutil.ReadAll(res.Body)
        return fmt.Errorf("failed to delete user in realm %s, received status code: %d, error message: %s", realm, res.StatusCode, string(errorResponse))
    }

    return nil
}



// GetUserIDByUsername retrieves the userID based on the provided username
func (kc *KeycloakClient) GetUserIDByUsername(realm, username string) (string, error) {
    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, realm)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer "+kc.Token)

    res, err := kc.Client.Do(req)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        errorResponse, _ := ioutil.ReadAll(res.Body)
        return "", fmt.Errorf("failed to retrieve users from realm %s, received status code: %d, error message: %s", realm, res.StatusCode, string(errorResponse))
    }

    var users []User
    err = json.NewDecoder(res.Body).Decode(&users)
    if err != nil {
        return "", err
    }

    for _, user := range users {
        if user.Username == username {
            return user.ID, nil
        }
    }

    return "", fmt.Errorf("user with username %s not found in realm %s", username, realm)
}

