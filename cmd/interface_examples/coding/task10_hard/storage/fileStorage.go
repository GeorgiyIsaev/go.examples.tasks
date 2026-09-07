package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type FileStorage struct {
	file *os.File
	mu   sync.Mutex // для защиты concurrent доступа если потребуется
}

// NewFileStorage открывает файл для чтения/записи (создаёт, если не существует).
func NewFileStorage(path string) (*FileStorage, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStorage{file: f}, nil
}

// readAllMap читает весь файл и парсит JSON в map.
func (fs *FileStorage) readAllMap() (map[string]interface{}, error) {
	// Сбрасываем указатель в начало
	if _, err := fs.file.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	data, err := io.ReadAll(fs.file)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return make(map[string]interface{}), nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		// Если файл повреждён, возвращаем пустую карту (или можно вернуть ошибку, но по условию лучше считать пустым)
		return make(map[string]interface{}), nil
	}
	return m, nil
}

// writeAllMap записывает map в файл, перезаписывая его целиком.
func (fs *FileStorage) writeAllMap(m map[string]interface{}) error {
	// Сбрасываем указатель и усекаем файл
	if err := fs.file.Truncate(0); err != nil {
		return err
	}
	if _, err := fs.file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	_, err = fs.file.Write(data)
	return err
}

func (fs *FileStorage) Get(key string) (interface{}, error) {
	m, err := fs.readAllMap()
	if err != nil {
		return nil, err
	}
	val, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return val, nil
}

func (fs *FileStorage) Set(key string, value interface{}) error {
	m, err := fs.readAllMap()
	if err != nil {
		return err
	}
	m[key] = value
	return fs.writeAllMap(m)
}

// Close реализует io.Closer.
func (fs *FileStorage) Close() error {
	return fs.file.Close()
}
