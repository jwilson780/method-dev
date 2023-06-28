package credentials

import (
	"encoding/json"
	"io"
	"os"
)

type Credentials struct {
	Username    string `json:"username"`
	OauthToken  string `json:"oauth_token"`
	ChannelName string `json:"channel_name"`
}

// LoadCredentials loads credentials from credentials.json
func LoadCredentials(path string) (*Credentials, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var creds Credentials
	err = json.Unmarshal(bytes, &creds)
	if err != nil {
		return nil, err
	}
	return &creds, nil
}
