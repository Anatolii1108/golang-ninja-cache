package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	Get(key string) any
	Set(key string, value any)
	Delete(key string)
}

type MemoryCache struct {
	mu   *sync.RWMutex
	data map[string]cacheItem
}

type cacheItem struct {
	value       any
	whenExpired time.Time
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]cacheItem), // What is the difference between make(map[string]any) and map[string]any{} ?
		mu:   new(sync.RWMutex),
	}
}

func (c *MemoryCache) Set(key string, value any, defaultItemLifetime ...time.Duration) error {
	if err := validateKey(key); err != nil {
		return err
	}

	if err := validateValue(value); err != nil {
		return err
	}

	lifetime := defaultLifetime
	if len(defaultItemLifetime) > 0 {
		lifetime = defaultItemLifetime[0]
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheItem{
		value:       value,
		whenExpired: getExpiration(lifetime),
	}

	return nil
}

func (c *MemoryCache) Get(key string) (any, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok || isExpired(item) {
		return nil, fmt.Errorf("key %q isn't exists or value is expired", key)
	}

	return item.value, nil
}

func (c *MemoryCache) Delete(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)

	return nil
}
