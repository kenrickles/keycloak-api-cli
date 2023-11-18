package keycloak

import (
    "time"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"

	"keycloak-api-cli/pkg/config"
)

// KeycloakClient struct holds data necessary to communicate with the Keycloak server
type KeycloakClient struct {
    BaseURL    string
    Client     *http.Client
    Token      string
    AuthRealm  string
    RealmToEdit string
}

// NewClient creates a new instance of KeycloakClient
func NewClient(cfg config.KeycloakConfig) *KeycloakClient {
	return &KeycloakClient{
        BaseURL: cfg.URL,
        Client: &http.Client{
            Timeout: time.Second * 30,
        },
        AuthRealm: cfg.AuthRealm,
        RealmToEdit: cfg.RealmToEdit,
    }
}

// Authenticate method for KeycloakClient to authenticate with the Keycloak server
func (kc *KeycloakClient) Authenticate(cfg config.KeycloakConfig) error {
    data := url.Values{}
    data.Set("client_id", cfg.ClientID)
    data.Set("grant_type", "password")
    data.Set("username", cfg.Username)
    data.Set("password", cfg.Password)

    // Include client_secret only if it's a confidential client
    if cfg.ClientSecret != "" {
        data.Set("client_secret", cfg.ClientSecret)
        data.Set("grant_type", "client_credentials")
    }

    // Construct the request URL
    requestURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", kc.BaseURL, kc.AuthRealm)

    // Create a new POST request
    req, err := http.NewRequest("POST", requestURL, strings.NewReader(data.Encode()))
    if err != nil {
        return fmt.Errorf("error creating request: %v", err)
    }
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    // Send the request
    resp, err := kc.Client.Do(req)
    if err != nil {
        return fmt.Errorf("error requesting token: %v", err)
    }
    defer resp.Body.Close()

    // Read the response body
    responseBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("error reading token response: %v", err)
    }

    // Check for non-200 HTTP response
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("received non-200 response status: %d - %s", resp.StatusCode, string(responseBody))
    }

    // Decode JSON response into a struct
    var tokenResponse struct {
        AccessToken string `json:"access_token"`
    }
    if err := json.Unmarshal(responseBody, &tokenResponse); err != nil {
        return fmt.Errorf("error decoding token response: %v", err)
    }

    kc.Token = tokenResponse.AccessToken
    return nil
}
