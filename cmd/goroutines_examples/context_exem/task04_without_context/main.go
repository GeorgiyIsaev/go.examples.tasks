package main

import (
	"context"
	"fmt"
	"time"
)

// ВОРКЕР без контекста
func stubbornWorker() {
	i := 0
	for {
		time.Sleep(80 * time.Millisecond)
		i++
		fmt.Printf("  упрямый воркер: тик %d (контекст я не слушаю)\n", i)
	}
}

// Пример: упрямый воркер (игнорирует контекст)
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// этот воркер НЕ получает ctx — он его просто не видит
	go stubbornWorker()

	// отменяем через 200 мс
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("  [main] вызываю cancel()")
		cancel()
	}()
	time.Sleep(500 * time.Millisecond)
	// ждём 500 мс и смотрим: отмена была, а воркер всё тикает
	select {
	case <-ctx.Done():
		fmt.Println("  [main] ctx.Done() сработал:", ctx.Err())
	case <-time.After(500 * time.Millisecond):
	}

	fmt.Println("  [main] смотрю: упрямый воркер всё ещё работает")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  [main] выхожу — процесс завершится и убьёт горутину вместе с собой")
}
