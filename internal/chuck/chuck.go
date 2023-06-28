package chuck

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}

// GetJoke retrieves a joke from the Chuck Norris API
func GetJoke(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", errors.New(fmt.Sprint("Failed to retrieve joke:", err.Error()))
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "", errors.New(fmt.Sprint("Failed to parse joke:", err.Error()))
	}
	return joke.Value, nil
}
