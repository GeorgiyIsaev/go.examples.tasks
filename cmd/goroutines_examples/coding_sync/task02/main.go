package main

import (
	"fmt"
	"sync"
)

//2. Потокобезопасный кэш на RWMutex
//   Задача: реализовать generic-кэш Cache[K comparable, V any]
//   с методами Get, Set, Delete.
//   Чтения должны идти параллельно, запись — эксклюзивно.
//   Тест: параллельные Get и Set, запуск с -race.

type Cache[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.m[key]
	return v, ok
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{m: make(map[K]V)}
}

func main() {
	cache := NewCache[int, string]()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(i, fmt.Sprintf("Номер %d", i))
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(cache.Get(i))
		}()
	}
	wg.Wait()

}
