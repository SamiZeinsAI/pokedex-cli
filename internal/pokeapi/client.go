package pokeapi

import (
	"net/http"
	"time"

	"github.com/SamiZeinsAI/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
