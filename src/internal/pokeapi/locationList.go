package pokeapi

import (
	"encoding/json"
	"errors"
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

func (c *Client) ExploreLocation(area string) (PokerAreaEncounters, error) {
	if area == "" {
		return PokerAreaEncounters{}, errors.New("Specify area to explore")
	}
	url := baseUrl + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		fmt.Printf("fetching %v from cache \n", url)
		encountersRes := PokerAreaEncounters{}
		err := json.Unmarshal(val, &encountersRes)
		if err != nil {
			return PokerAreaEncounters{}, err
		}

		return encountersRes, nil
	}

	fmt.Printf("fetching %v from pokeapi \n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokerAreaEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokerAreaEncounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokerAreaEncounters{}, err
	}

	encountersRes := PokerAreaEncounters{}
	err = json.Unmarshal(data, &encountersRes)
	if err != nil {
		return PokerAreaEncounters{}, err
	}

	c.cache.Add(url, data)
	return encountersRes, nil

}
