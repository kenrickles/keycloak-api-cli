package keycloak

import (
    "encoding/json"
	"bytes"
    "fmt"
    "net/http"
)

// Realm represents a Keycloak realm
type Realm struct {
    ID          string `json:"id"`
    Realm       string `json:"realm"`
    DisplayName string `json:"displayName"`
}

type RealmDetails struct {
    Realm       string `json:"realm"`
    DisplayName string `json:"displayName"`
}


// ListRealms lists all realms in Keycloak
func (kc *KeycloakClient) ListRealms() ([]Realm, error) {
    url := fmt.Sprintf("%s/admin/realms", kc.BaseURL)
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

    var realms []Realm
    err = json.NewDecoder(res.Body).Decode(&realms)
    if err != nil {
        return nil, err
    }

    return realms, nil
}

func (kc *KeycloakClient) CreateRealm(realmName string) error {
	details := RealmDetails{
        Realm: realmName,
    }
    url := fmt.Sprintf("%s/admin/realms", kc.BaseURL)

    // Encode the realm details as JSON for the request body
    jsonBody, err := json.Marshal(details)
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
        return fmt.Errorf("failed to create realm, received status code: %d", res.StatusCode)
    }
	
    return nil
}