package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationAreasResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, ok := c.cache.Get(url)
	if !ok {
		var err error
		body, err = c.fetch(url)
		if err != nil {
			return LocationAreasResp{}, err
		}
		c.cache.Add(url, body)
	}

	var data LocationAreasResp
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreasResp{}, err
	}
	return data, nil
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name
	body, ok := c.cache.Get(url)
	if !ok {
		var err error
		body, err = c.fetch(url)
		if err != nil {
			return LocationArea{}, err
		}
		c.cache.Add(url, body)
	}
	var data LocationArea
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationArea{}, err
	}
	return data, nil
}

func (c *Client) fetch(url string) ([]byte, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	return body, nil
}
