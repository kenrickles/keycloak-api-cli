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
func (kc *KeycloakClient) CreateUser(username, email, password string) error {
    userDetails := struct {
        Username   string `json:"username"`
        Email      string `json:"email"`
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
                Type:  "password",
                Value: password,
            },
        },
    }

    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, kc.RealmToEdit)

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
        return fmt.Errorf("failed to create user, received status code: %d, error message: %s", res.StatusCode, errorResponse)
    }

    return nil
}

// ListUsers lists all users in Keycloak
func (kc *KeycloakClient) ListUsers() ([]User, error) {
    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, kc.RealmToEdit)
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

func (kc *KeycloakClient) DeleteUser(userID string) error {
    url := fmt.Sprintf("%s/admin/realms/%s/users/%s", kc.BaseURL, kc.RealmToEdit, userID)

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

    // Check for successful status code, typically 204 for successful deletion
    if res.StatusCode != http.StatusNoContent {
        errorResponse, _ := ioutil.ReadAll(res.Body)
        return fmt.Errorf("failed to delete user, received status code: %d, error message: %s", res.StatusCode, errorResponse)
    }

    return nil
}


// GetUserIDByUsername retrieves the userID based on the provided username
func (kc *KeycloakClient) GetUserIDByUsername(username string) (string, error) {
    url := fmt.Sprintf("%s/admin/realms/%s/users", kc.BaseURL, kc.RealmToEdit)

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
        return "", fmt.Errorf("failed to retrieve users, received status code: %d, error message: %s", res.StatusCode, errorResponse)
    }

    // Read the response body and parse it as an array of User objects
    var users []User
    err = json.NewDecoder(res.Body).Decode(&users)
    if err != nil {
        return "", err
    }

    // Search for the user with the matching username
    for _, user := range users {
        if user.Username == username {
            return user.ID, nil
        }
    }

    // If no matching user is found, return an error
    return "", fmt.Errorf("user with username %s not found in %s realm", username, kc.RealmToEdit)
}
