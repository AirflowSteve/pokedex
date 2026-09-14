package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]CacheEntry
	mu    sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		timer := time.Now()
		for entry, value := range c.cache {
			if timer.Sub(value.createdAt) > interval {
				c.mu.Lock()
				delete(c.cache, entry)
				c.mu.Unlock()
			}
		}
	}
}

func NewCache(interval time.Duration) {
	cache := Cache{}
	go cache.reapLoop(5 * time.Second)
}
