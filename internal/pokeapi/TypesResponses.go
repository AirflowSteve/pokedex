package pokeapi

type pokeAPIResponse struct {
	Count       int            `json:"count"`
	NextURL     string         `json:"next"`
	PreviousURL string         `json:"previous"`
	Results     []pokeLocation `json:"results"`
}

type pokeLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
