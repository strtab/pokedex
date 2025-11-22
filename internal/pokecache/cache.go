package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (cache *Cache) {
	cache = &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	cache.reapLoop()
	return
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, exists := cache.entries[key]
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) reapLoop() {
	go func() {
		ticker := time.NewTicker(cache.interval)
		defer ticker.Stop()

		for range ticker.C {
			cache.mu.Lock()

			for key, val := range cache.entries {
				if time.Since(val.createdAt) > cache.interval {
					delete(cache.entries, key)
				}
			}

			cache.mu.Unlock()
		}
	}()
}
