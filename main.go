package main

import (
	"time"

	pockecache "github.com/arawwad/pokedexcli/internal/pokecache"
)

type serviceConfig struct {
	next     string
	previous string
}

var config = serviceConfig{
	next: "https://pokeapi.co/api/v2/location-area?offset=0",
}

var cache = pockecache.NewCache(5 * time.Minute)

func main() {
	defer cache.Stop()
	startRepl()
}
