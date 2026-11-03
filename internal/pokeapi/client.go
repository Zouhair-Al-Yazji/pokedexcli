// Package pokeapi provides a client to interact with the PokeAPI service,
// allowing the pokedexcli to fetch location areas and pokemon data.
package pokeapi

import (
	"net/http"
	"time"

	"github.com/Zouhair-Al-Yazji/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}
