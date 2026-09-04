package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func (c *Cache) Store(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Load(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func main() {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	// Сохранение данных
	cache.Store("name", "Alice")
	cache.Store("age", 25)

	// Загрузка с проверкой существования и типа
	name1, ok1 := cache.Load("age").(string)
	if !ok1 {
		fmt.Println("Name is not a string or not found")
	} else {
		fmt.Println("Name:", name1)
	}
	// Загрузка с проверкой существования и типа
	name2, ok2 := cache.Load("name").(string)
	if !ok2 {
		fmt.Println("Name is not a string or not found")
	} else {
		fmt.Println("Name:", name2)
	}

	// Загрузка с приведением без проверки (может вызвать panic)
	age := cache.Load("age").(int)
	fmt.Println("Age:", age)

	// Пример загрузки отсутствующего ключа
	height, ok := cache.Load("height").(float64)
	if !ok {
		fmt.Println("Height not found or not a float64")
	} else {
		fmt.Println("Height:", height)
	}

	// Альтернативная проверка через nil (если значение может быть nil)
	heightVal := cache.Load("height")
	if heightVal == nil {
		fmt.Println("Height not found")
	} else {
		fmt.Println("Height:", heightVal)
	}

	//Паника нет проверки наличия, нет проверки типа, нет проверки nil
	//width := cache.Load("width").(float64)
	//fmt.Println("Width:", width)

}
