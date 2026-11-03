// Package pokecache provides a small in-memory cache with automatic expiry.
// Entries older than the configured interval are removed by a background
// reap loop. It is safe for concurrent use.
package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	cutoff := time.Now().Add(-c.interval)
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
