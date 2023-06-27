package chuck

import (
	"encoding/json"
	"io"
	"net/http"
)

const ChuckApi string = "https://api.chucknorris.io/jokes/random"
const ChuckMessage string = "!chucknorris"

type Joke struct {
	Value string `json:"value"`
}

func GetJoke() string {
	resp, err := http.Get(ChuckApi)
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
