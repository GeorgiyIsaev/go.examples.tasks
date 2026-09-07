package storage

import "fmt"

type MapStorage struct {
	data map[string]interface{}
}

func NewMapStorage() *MapStorage {
	return &MapStorage{data: make(map[string]interface{})}
}

func (m *MapStorage) Get(key string) (interface{}, error) {
	val, ok := m.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return val, nil
}

func (m *MapStorage) Set(key string, value interface{}) error {
	m.data[key] = value
	return nil
}
