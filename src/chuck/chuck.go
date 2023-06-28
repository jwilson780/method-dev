package chuck

import (
	"encoding/json"
	"io"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}

// GetJoke retrieves a joke from the Chuck Norris API
func GetJoke(apiURL string) string {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "Failed to retrieve joke: " + err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Failed to parse joke: " + err.Error()
	}
	return joke.Value
}
