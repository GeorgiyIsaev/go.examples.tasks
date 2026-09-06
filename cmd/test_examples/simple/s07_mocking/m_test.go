package s07_mocking

import (
	"errors"
	"testing"
)

//Библиотеки Моков
//GoMock — генерация моков из интерфейсов
//Testify — набор утилит для тестирования (assertions + mocking)
//gock — HTTP-моки без реальных запросов

// Определяем интерфейс
type Database interface {
	Get(id int) (string, error)
}

// Реальная реализация
type RealDB struct{}

func (db RealDB) Get(id int) (string, error) {
	// обращение к БД
}

// Мок для тестов
type MockDB struct {
	data map[int]string
}

func (db MockDB) Get(id int) (string, error) {
	if val, ok := db.data[id]; ok {
		return val, nil
	}
	return "", errors.New("not found")
}

// Тест с моком
func TestService(t *testing.T) {
	mockDB := MockDB{data: map[int]string{1: "test"}}
	service := NewService(mockDB)
	// тестирование...
}
