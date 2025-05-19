package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		CacheEntries: map[string]cacheEntry{},
		Mutex:        sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.CacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	cache, exists := c.CacheEntries[key]
	if exists {
		return cache.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Mutex.Lock()
		now := time.Now()
		for key, val := range c.CacheEntries {
			if now.Sub(val.createdAt) > interval {
				delete(c.CacheEntries, key)
			}
		}
		c.Mutex.Unlock()
	}
}
