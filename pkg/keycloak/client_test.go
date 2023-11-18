package keycloak

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "keycloak-api-cli/pkg/config"
)

func TestAuthenticate(t *testing.T) {
    // Mock server to simulate Keycloak response
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"access_token": "test-token"}`))
    }))
    defer server.Close()

    cfg := config.KeycloakConfig{
        URL:          server.URL, // Use mock server URL
        ClientID:     "test-id",
        ClientSecret: "secret",
        Username:     "user",
        Password:     "pass",
        AuthRealm:    "realm-test",
    }

    client := NewClient(cfg)
    err := client.Authenticate(cfg)
    if err != nil {
        t.Errorf("Authenticate failed: %v", err)
    }
    if client.Token != "test-token" {
        t.Errorf("Expected token 'test-token', got '%s'", client.Token)
    }
}
