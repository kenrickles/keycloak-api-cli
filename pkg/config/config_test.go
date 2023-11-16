package config

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
    "reflect"

    "github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
    // Create a temporary directory
    tempDir, err := ioutil.TempDir("", "config")
    if err != nil {
        t.Fatalf("Failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir) // Clean up

    // Create a temporary config file within the temp directory
    tempFilePath := filepath.Join(tempDir, "config.yaml")
    tempConfig := []byte("url: http://example.com\nclient_id: test-id\nclient_secret: secret\nusername: user\npassword: pass\nrealm: realm-test")
    if err := ioutil.WriteFile(tempFilePath, tempConfig, 0666); err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }

    // Set Viper to use the temporary config file
    viper.Reset()
    viper.SetConfigName("config") 
    viper.SetConfigType("yaml")
    viper.AddConfigPath(tempDir)       
    viper.AutomaticEnv()

    // Call the LoadConfig function
    gotConfig, err := LoadConfig()
    if err != nil {
        t.Fatalf("LoadConfig() error = %v", err)
    }

    // Expected configuration
    expectedConfig := KeycloakConfig{
        URL:          "http://example.com",
        ClientID:     "test-id",
        ClientSecret: "secret",
        Username:     "user",
        Password:     "pass",
        Realm:        "realm-test",
    }

    // Compare the result with the expected configuration
    if !reflect.DeepEqual(gotConfig, expectedConfig) {
        t.Errorf("LoadConfig() = %v, want %v", gotConfig, expectedConfig)
    }
}
