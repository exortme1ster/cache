package cache

type Cache struct {
	memoryCache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		memoryCache: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.memoryCache[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c.memoryCache[key]
}

func (c Cache) Delete(key string) {
	delete(c.memoryCache, key)
}
