package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (PokeLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Printf("fetching %v from cache \n", url)
		locationsRes := PokeLocations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return PokeLocations{}, err
		}

		return locationsRes, nil
	}

	fmt.Printf("fetching %v from pokeapi \n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLocations{}, err
	}

	locationsRes := PokeLocations{}
	err = json.Unmarshal(data, &locationsRes)
	if err != nil {
		return PokeLocations{}, err
	}

	c.cache.Add(url, data)
	return locationsRes, nil
}
