package pokecache

import (
	"sync"
	"time"
)

// It's time to implement caching! This will make moving around the map feel a lot snappier. We'll be building a flexible caching system to help with performance in future steps.

// In our case, we'll be caching responses from the PokeAPI so that when we need that same data again, we can grab it from memory instead of making another network request.

// I[the instructor] used a Cache struct to hold a map[string]cacheEntry and a mutex to protect the map across goroutines.
type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.Mutex

	interval time.Duration
	stopCh   chan struct{}
	wg       sync.WaitGroup
}

// A cacheEntry should be a struct with two fields:
// -	createdAt - A time.Time that represents when the entry was created.
// -	val - A []byte that represents the raw data we're caching.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Feel free to add some logging that informs you in the command line when the cache is being used.

// You'll probably want to expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
		stopCh:   make(chan struct{}),
	}

	// todo: I don't understand this
	go func() {
		defer c.wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.reapLoop()
			case <-c.stopCh:
				return
			}
		}
	}()

	return c
}

// Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.cacheMap[key] = ce
}

// Create a cache.Get() method that gets an entry from the cache. It should take a key (a string) and return a []byte and a bool.
// The bool should be true if the entry was found and false if it wasn't.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}

	return entry.val, ok

}

// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval.
// This makes sure that the cache doesn't grow too large over time.
// For example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
// I used a time.Ticker to make this happen.
func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	cutoff := time.Now().Add(-c.interval)
	for k, v := range c.cacheMap {
		if v.createdAt.Before(cutoff) {
			delete(c.cacheMap, k)
		}
	}
}

// Maps are not thread-safe in Go.
// You should use a sync.Mutex to lock access to the map when you're adding, getting entries or reaping entries.
// It's unlikely that you'll have issues because reaping only happens every ~5 seconds, but it's still possible,
// so you should make your cache package safe for concurrent use.
