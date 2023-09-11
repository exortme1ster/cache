package cache

import (
	"sync"
	"time"
)

type item struct {
	value  interface{}
	expiry time.Time
}

type Cache struct {
	memoryCache map[string]item
	mu          *sync.RWMutex
	ticker      *time.Ticker
}

func New() *Cache {
	c := &Cache{
		ticker:      time.NewTicker(time.Second * 1),
		memoryCache: make(map[string]item),
		mu:          new(sync.RWMutex),
	}

	go c.cleanupExpiredItems()

	return c
}

func (c Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.memoryCache[key] = item{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	it, ok := c.memoryCache[key]
	if !ok {
		return nil
	}

	return it.value
}

func (c Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.memoryCache, key)
	c.mu.Unlock()
}

func (c Cache) cleanupExpiredItems() {
	for {
		select {
		case <-c.ticker.C:
			c.mu.Lock()
			for key, it := range c.memoryCache {
				if time.Now().After(it.expiry) {
					delete(c.memoryCache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
