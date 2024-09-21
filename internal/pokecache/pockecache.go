package pockecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]CacheEntry
	ttl   time.Duration
	done  chan bool
}

func NewCache(ttl time.Duration) *Cache {

	cache := Cache{
		items: make(map[string]CacheEntry),
		ttl:   ttl,
		done:  make(chan bool),
	}

	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	c.items[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheEntry, ok := c.items[key]

	if !ok {
		return nil, false
	}

	if time.Since(cacheEntry.createdAt) > c.ttl {
		return nil, false
	}

	return cacheEntry.val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()

			for key, value := range c.items {
				if time.Since(value.createdAt) > c.ttl {
					delete(c.items, key)
				}
			}

			c.mu.Unlock()

		case <-c.done:
			return
		}
	}

}

func (c *Cache) Stop() {
	c.done <- true
}
