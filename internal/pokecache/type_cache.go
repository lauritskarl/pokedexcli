package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheEntries map[string]cacheEntry
	Mutex        sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
