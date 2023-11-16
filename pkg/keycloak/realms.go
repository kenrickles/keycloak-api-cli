package keycloak

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Realm represents a Keycloak realm
type Realm struct {
    ID          string `json:"id"`
    Realm       string `json:"realm"`
    DisplayName string `json:"displayName"`
    // Add other fields as needed
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
