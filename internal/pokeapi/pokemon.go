package pokeapi

import "encoding/json"

type Pokemon struct {
	Name  string `json:"name"`
	Stats []struct {
		State struct {
			Name string `json:"name"`
		} `json:"state"`
		BaseState int `json:"base_state"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
	Height         int `json:"height"`
	Weight         int `json:"weight"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	body, ok := c.cache.Get(url)
	if !ok {
		var err error
		body, err = c.fetch(url)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, body)
	}
	var data Pokemon
	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, err
	}
	return data, nil
}
