package cache

type Cache interface {
	Get(key string) any
	Set(key string, value any)
	Delete(key string)
}

type MemoryCache struct {
	data map[string]any
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		data: make(map[string]any), // What is the difference between make(map[string]any) and map[string]any{} ?
	}
}

func (c *MemoryCache) Get(key string) any {
	return c.data[key]
}

func (c *MemoryCache) Set(key string, value any) {
	c.data[key] = value
}

func (c *MemoryCache) Delete(key string) {
	delete(c.data, key)
}
