package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	data     map[string]CacheEntry
	interval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{
		data:     make(map[string]CacheEntry),
		interval: duration,
	}
	cache.reaploop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.data[key] = cacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (c *Cache) reaploop() {
	ticker := time.NewTicker(c.interval)
	fmt.Println("Starting reaploop with interval:", c.interval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			c.mu.Lock()

			for key, val := range c.data {
				if time.Since(val.createdAt) > c.interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
