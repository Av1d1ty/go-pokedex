package pokekache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	data      []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        cache: make(map[string]cacheEntry),
    }
    go cache.reapLoop(interval)
    return cache
}

func (c *Cache) Set(key string, data []byte) {
    c.cache[key] = cacheEntry{
        data:      data,
        createdAt: time.Now(),
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    entry, ok := c.cache[key]
    if !ok {
        return nil, false
    }
    return entry.data, true
}

func (c *Cache) Delete(key string) {
    delete(c.cache, key)
}

func (c *Cache) reap(interval time.Duration) {
    for k, v := range c.cache {
        if time.Since(v.createdAt) > interval {
            c.Delete(k)
        }
    }
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C{
        c.reap(interval)
        time.Sleep(interval)
    }
}
