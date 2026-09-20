package main

import (
	"context"
	"fmt"
	"time"
)

//Пример: Отмена горутины через WithCancel
//cancel() закрывает канал ctx.Done().
//Все горутины, которые слушают этот канал в select,
//узнают об отмене.

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // сюда прилетит сигнал при cancel()
			fmt.Printf("worker %d: остановлен, причина: %v\n", id, ctx.Err())
			return
		default:
			// имитация полезной работы
			fmt.Printf("worker %d: работаю...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// Создаём контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // хорошая практика — всегда вызывать cancel, чтобы освободить ресурсы

	go worker(ctx, 1)
	go worker(ctx, 2)

	// Даём поработать 2 секунды
	time.Sleep(2 * time.Second)

	fmt.Println("main: посылаю сигнал отмены")
	cancel() // <- вот здесь все воркеры получают сигнал

	// Немного ждём, чтобы увидеть сообщения о завершении
	time.Sleep(200 * time.Millisecond)
	fmt.Println("main: выход")
}
