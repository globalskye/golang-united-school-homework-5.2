package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]Data
}

type Data struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{
		data: map[string]Data{},
	}

}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key1, value1 := range c.data {
		if key == key1 {

			if value1.deadline.After(time.Now()) {
				return value1.value, true
			} else {
				return value1.value, false
			}
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = Data{
		value:    value,
		deadline: time.Now().Add(time.Hour * 10),
	}

}

func (c *Cache) Keys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	keys := make([]string, 0, len(c.data))
	for key := range c.data {
		if c.data[key].deadline.After(time.Now()) {
			keys = append(keys, key)
		}

	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = Data{
		value:    value,
		deadline: deadline,
	}
}
