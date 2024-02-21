package pokekache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
        mux:   &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Set(key string, data []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		data:      data,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.Lock()
    defer c.mux.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Delete(key string) {
    c.mux.Lock()
    defer c.mux.Unlock()
	delete(c.cache, key)
}

func (c *Cache) reap(interval time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()
	for k, v := range c.cache {
		if time.Since(v.createdAt) > interval {
            delete(c.cache, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
		time.Sleep(interval)
	}
}
