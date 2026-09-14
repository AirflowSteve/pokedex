package pokeapi

import (
	"encoding/json"
	"net/http"
)

func RequestAndResponse(url string, client *http.Client) (pokeAPIResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeAPIResponse{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return pokeAPIResponse{}, err
	}
	defer res.Body.Close()

	var response pokeAPIResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return pokeAPIResponse{}, err
	}
	return response, nil

}
