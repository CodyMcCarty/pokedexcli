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
	mu       sync.Mutex
	store    map[string]cacheEntry
	interval time.Duration
	stopCh   chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		store:    make(map[string]cacheEntry),
		interval: interval,
		stopCh:   make(chan struct{}),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	cp := make([]byte, len(val))
	copy(cp, val)

	c.mu.Lock()
	c.store[key] = cacheEntry{createdAt: time.Now(), val: cp}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry,
}
