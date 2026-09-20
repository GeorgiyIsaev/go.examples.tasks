package main

import (
	"context"
	"fmt"
	"time"
)

//Пример 2. Таймаут через WithTimeout
//кейс: «если запрос не уложился в N секунд — бросаем».

//При	case <-time.After(1 * time.Second) - сработает
//При 	case <-time.After(5 * time.Second) - будет ошибка ожидания

// Долгая операция, которая умеет уважать контекст
func longOperation(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second): // как будто долгий запрос
		return "результат", nil
	case <-ctx.Done(): // таймаут или отмена
		return "", ctx.Err()
	}
}

func main() {
	// Контекст с таймаутом 1 секунда
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := longOperation(ctx)
	if err != nil {
		fmt.Println("ошибка:", err) // context deadline exceeded
		return
	}
	fmt.Println("успех:", result)
}
