package pokeapi

import (
	"net/http"
	"time"

	"github.com/lauritskarl/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(5 * time.Second),
	}
}
