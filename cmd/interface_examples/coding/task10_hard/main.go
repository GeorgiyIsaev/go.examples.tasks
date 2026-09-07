package main

import (
	"fmt"
	"io"
	"log"

	"go.examples.tasks/cmd/interface_examples/coding/task10_hard/decorator"
)

// Проверяем есть ли io.Closer, если да логируем
func CloseIfPossible(v interface{}) {
	if closer, ok := v.(io.Closer); ok {
		if err := closer.Close(); err != nil {
			log.Printf("Close error: %v", err)
		} else {
			log.Println("Closed successfully")
		}
	} else {
		// Ничего не делаем
	}
}

func main() {
	// 1. Строим цепочку: metrics(logging(file("data.json")))
	fileStore, err := NewStorage("file", "data.json")
	if err != nil {
		log.Fatal(err)
	}
	logStore, err := NewStorage("logging", fileStore)
	if err != nil {
		log.Fatal(err)
	}
	metricStore, err := NewStorage("metrics", logStore)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Выполняем операции
	storage := metricStore // внешняя обёртка
	_ = storage.Set("name", "Alice")
	_ = storage.Set("age", 30)
	val, _ := storage.Get("name")
	fmt.Printf("Get name: %v\n", val)
	_, _ = storage.Get("missing") // вызовет ошибку

	// 3. Извлекаем Metrics
	if metrics, ok := storage.(decorator.Metrics); ok {
		fmt.Printf("GetCalls: %d, SetCalls: %d, GetErrors: %d, SetErrors: %d\n",
			metrics.GetCalls(), metrics.SetCalls(), metrics.GetErrors(), metrics.SetErrors())
	} else {
		fmt.Println("Metrics not available")
	}

	// 4. Закрываем все ресурсы
	// Передаём внешнюю обёртку – она закроет внутренние через делегирование
	CloseIfPossible(storage)
}
