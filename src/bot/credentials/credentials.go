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

func LoadCredentials() (*Credentials, error) {
	file, err := os.Open("credentials.json")
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
