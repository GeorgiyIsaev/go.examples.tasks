package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализуйте потокобезопасный кэш с временем жизни (TTL).

type item struct {
	value      interface{}
	expiration int64
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]item
	stop  chan struct{}
}

func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]item),
		stop:  make(chan struct{}),
	}
	go cache.clenupLoop()
	return cache
}

// сохраняет значение с TTL.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = item{
		value:      value,
		expiration: time.Now().Add(ttl).Unix(),
	}
}

// возвращает значение, если оно есть и не истекло.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	if time.Now().UnixNano() > item.expiration {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
		return nil, false
	}
	return item.value, true

}

func (c *Cache) clenupLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.RLock()
			now := time.Now().UnixNano()
			for key, item := range c.items {
				if now > item.expiration {
					delete(c.items, key)
				}
			}
			c.mu.RUnlock()
		case <-c.stop:
			return
		}
	}
}

func (c *Cache) Close() {
	close(c.stop)
}

func main() {
	cache := NewCache()
	cache.Set("hello", "world", 2*time.Second)
	v, ok := cache.Get("key")
	fmt.Println("v", v, ok)
	time.Sleep(3 * time.Second)
	v, ok = cache.Get("key")
	fmt.Println("v", v, ok)

}
