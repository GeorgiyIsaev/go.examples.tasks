package decorator

import (
	"io"
	"log"
	"time"

	"go.examples.tasks/cmd/interface_examples/coding/task10_hard/storage"
)

type LoggingStorage struct {
	inner storage.Storage
}

func NewLoggingStorage(inner storage.Storage) *LoggingStorage {
	return &LoggingStorage{inner: inner}
}

func (ls *LoggingStorage) Get(key string) (interface{}, error) {
	start := time.Now()
	val, err := ls.inner.Get(key)
	log.Printf("[LOG] Get key=%s, took=%v, err=%v", key, time.Since(start), err)
	return val, err
}

func (ls *LoggingStorage) Set(key string, value interface{}) error {
	start := time.Now()
	err := ls.inner.Set(key, value)
	log.Printf("[LOG] Set key=%s, took=%v, err=%v", key, time.Since(start), err)
	return err
}

// Close закрывает внутреннее хранилище, если оно поддерживает io.Closer.
func (ls *LoggingStorage) Close() error {
	if closer, ok := ls.inner.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
