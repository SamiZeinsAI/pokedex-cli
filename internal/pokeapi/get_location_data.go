package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(name string) (LocationData, error) {
	url := baseUrl + "location-area/" + name
	fmt.Println(url)
	locationData := LocationData{}
	if val, exists := c.cache.Get(url); exists {
		fmt.Println("used cache")
		err := json.Unmarshal(val, &locationData)
		if err != nil {
			return LocationData{}, nil
		}
		return locationData, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return LocationData{}, err
	}
	fmt.Println("didn't use cache")
	c.cache.Add(url, body)
	return locationData, nil
}
