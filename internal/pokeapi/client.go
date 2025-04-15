package pokeapi

import (
	"net/http"
	"time"

	"github.com/darginmathi/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	pokeDex    map[string]PokemonInfo
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeDex: make(map[string]PokemonInfo),
	}
}
