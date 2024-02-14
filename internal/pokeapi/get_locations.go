package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationArea, error) {

	url := baseUrl + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}

		defer resp.Body.Close()
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(url, body)
	}
	locArea := LocationArea{}
	err := json.Unmarshal(body, &locArea)
	if err != nil {
		fmt.Println("error at unmarshal")
		return LocationArea{}, err
	}
	return locArea, nil
}
