package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]CacheEntry
	mu    sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Cache[key] = CacheEntry{
		createdAt: time.Now(),
		Val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key == "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20" {
		key = "https://pokeapi.co/api/v2/location-area/"
	}

	entry, ok := c.Cache[key]
	if !ok {
		return nil, ok
	}
	return entry.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	c.mu.Lock()
	defer c.mu.Unlock()
	for range ticker.C {
		timer := time.Now()

		for entry, value := range c.Cache {
			if timer.Sub(value.createdAt) > interval {

				delete(c.Cache, entry)

			}
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Cache: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}
