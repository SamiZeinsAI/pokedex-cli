package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Duration(interval))
	for t := range ticker.C {
		c.reap(t.Add(-1 * interval))
	}
}
func (c *Cache) reap(cutoff time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
