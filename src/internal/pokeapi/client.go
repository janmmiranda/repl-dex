package pokeapi

import (
	"net/http"
	"time"

	"github.com/janmmiranda/repl-dex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Pokecache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewPokecache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
