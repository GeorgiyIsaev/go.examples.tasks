package decorator

import (
	"io"
	"sync"

	"go.examples.tasks/cmd/interface_examples/coding/task10_hard/storage"
)

type MetricsStorage struct {
	inner     storage.Storage
	getCalls  int
	setCalls  int
	getErrors int
	setErrors int
	mu        sync.Mutex // для потокобезопасности счетчиков
}

func NewMetricsStorage(inner storage.Storage) *MetricsStorage {
	return &MetricsStorage{inner: inner}
}

func (ms *MetricsStorage) Get(key string) (interface{}, error) {
	ms.mu.Lock()
	ms.getCalls++
	ms.mu.Unlock()
	val, err := ms.inner.Get(key)
	if err != nil {
		ms.mu.Lock()
		ms.getErrors++
		ms.mu.Unlock()
	}
	return val, err
}

func (ms *MetricsStorage) Set(key string, value interface{}) error {
	ms.mu.Lock()
	ms.setCalls++
	ms.mu.Unlock()
	err := ms.inner.Set(key, value)
	if err != nil {
		ms.mu.Lock()
		ms.setErrors++
		ms.mu.Unlock()
	}
	return err
}

// Реализация Metrics
func (ms *MetricsStorage) GetCalls() int {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.getCalls
}
func (ms *MetricsStorage) SetCalls() int {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.setCalls
}
func (ms *MetricsStorage) GetErrors() int {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.getErrors
}
func (ms *MetricsStorage) SetErrors() int {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.setErrors
}

// Close делегирует внутреннему.
func (ms *MetricsStorage) Close() error {
	if closer, ok := ms.inner.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
