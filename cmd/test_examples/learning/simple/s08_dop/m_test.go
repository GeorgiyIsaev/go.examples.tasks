package s08_dop

import "testing"

// TestParallel паралельные тесты
func TestParallel(t *testing.T) {
	t.Parallel() // Этот тест будет выполняться параллельно

	// тело теста
}

// TestLongRunning пропуск долгих тестов
func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	// долгий тест
}

// TestPanic - Проверка на панику
func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got none")
		}
	}()
	// код, который должен вызвать панику
}
