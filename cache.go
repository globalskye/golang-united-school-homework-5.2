package cache

import (
	
	"time"
)

type Cache struct {
	Key, Value string
	Deadline   time.Time
}


func NewCache(key string, value string, deadline time.Time) Cache {
	return Cache{Key: key, Value: value, Deadline: deadline}
}

func (c Cache) Get(key string) (string, bool) {
	return "", false

}
func (c *Cache) Put(key, value string) {
	NewCache(key, value, time.Now())

}

func (c Cache) Keys() []string {
	return nil
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	NewCache(key, value, deadline)
}