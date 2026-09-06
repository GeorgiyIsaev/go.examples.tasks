package s05_testMain

import (
	"os"
	"testing"
)

// TestMain позволяет выполнять подготовку и очистку перед запуском всех тестов в пакете
func TestMain(m *testing.M) {
	// Настройка перед тестами
	setup()

	// Запуск тестов
	code := m.Run()

	// Очистка после тестов
	teardown()

	os.Exit(code)
}
