package main

import (
	"fmt"
	"sync"
	"time"
)

type cacheValue struct {
	value      interface{}
	expiration time.Time
}

type ConcurrentCache struct {
	mu    sync.RWMutex
	cache map[string]cacheValue
}

func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		cache: make(map[string]cacheValue),
	}
}

func (c *ConcurrentCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheValue{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

func (c *ConcurrentCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.cache[key]
	if !found {
		return nil, false
	}

	if time.Now().After(item.expiration) {
		// If the data has expired, remove it from the cache and return not found.
		delete(c.cache, key)
		return nil, false
	}

	return item.value, true
}

func main() {
	cache := NewConcurrentCache()

	// Set some data in the cache with an expiration of 5 seconds
	cache.Set("key1", "data1", 5*time.Second)
	cache.Set("key2", "data2", 10*time.Second)

	// Retrieve data from the cache
	if data, found := cache.Get("key1"); found {
		fmt.Println("Data for key1:", data)
	} else {
		fmt.Println("Data for key1 not found.")
	}

	if data, found := cache.Get("key2"); found {
		fmt.Println("Data for key2:", data)
	} else {
		fmt.Println("Data for key2 not found.")
	}

	// Wait for some time to let the data with key1 expire
	time.Sleep(6 * time.Second)

	// Try to retrieve the expired data (it should not be found)
	if data, found := cache.Get("key1"); found {
		fmt.Println("Data for key1:", data)
	} else {
		fmt.Println("Data for key1 not found.")
	}
}
