package credentials

import (
	"encoding/json"
	"os"
	"testing"
)

func TestLoadCredentials(t *testing.T) {
	// create a temp file
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	// cleanup after test
	defer os.Remove(tmpfile.Name()) // clean up

	// create a Credentials instance
	creds := &Credentials{
		Username:    "testUser",
		OauthToken:  "oauth:testToken",
		ChannelName: "testChannel",
	}

	// marshal the credentials to JSON
	credsJSON, err := json.Marshal(creds)
	if err != nil {
		t.Fatalf("Failed to marshal credentials: %v", err)
	}

	// write the JSON to the temp file
	if _, err := tmpfile.Write(credsJSON); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// close the file
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// attempt to load credentials from the temp file
	loadedCreds, err := LoadCredentials(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to load credentials: %v", err)
	}

	// verify that the loaded credentials match the original
	if loadedCreds.Username != creds.Username || loadedCreds.OauthToken != creds.OauthToken || loadedCreds.ChannelName != creds.ChannelName {
		t.Fatalf("Loaded credentials do not match the original ones")
	}
}