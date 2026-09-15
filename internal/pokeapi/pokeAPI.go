package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func Request(url string, client http.Client) ([]byte, error) {
	// println("\n" + "Making a request" + "\n")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Decipher(data []byte) (PokeAPIResponse, error) {
	var response PokeAPIResponse

	err := json.Unmarshal(data, &response)
	if err != nil {
		return PokeAPIResponse{}, err
	}
	return response, nil
}
