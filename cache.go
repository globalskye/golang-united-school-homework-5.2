package main

import (
	"time"
)

type Cache struct {
	data map[string]Data
}

type Data struct {
	value    string
	deadline time.Time
}

func main() {
	c := NewCache()
	c.Put("Ключ", "Значение")
}

func NewCache() Cache {
	return Cache{
		data: map[string]Data{},
	}

}

func (c *Cache) Get(key string) (string, bool) {
	for key1, value1 := range c.data {
		if key == key1 {
			if value1.deadline.After(time.Now()) {
				return key1, true
			} else {
				return key1, false
			}
		}
	}
	return "", false
}
func (c *Cache) Put(key, value string) {
	c.data[key] = Data{
		value:    value,
		deadline: time.Now().Add(time.Minute),
	}
}

func (c *Cache) Keys() []string {
	keys := []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = Data{value, deadline}
}
