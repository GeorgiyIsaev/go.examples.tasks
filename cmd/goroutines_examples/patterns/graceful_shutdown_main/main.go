package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	// ctx отменится при получении SIGINT или SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	var counter int64

	// Горутина-инкрементатор
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return // сигнал получен — выходим из цикла
			default:
				atomic.AddInt64(&counter, 1)
			}
		}
	}()

	// Ждём сигнал
	<-ctx.Done()
	fmt.Println("\n[main] получен сигнал, дожидаемся горутины...")
	<-done // дожидаемся завершения горутины

	// Читаем итоговое значение и сохраняем в файл
	val := atomic.LoadInt64(&counter)
	content := fmt.Sprintf("counter=%d\ndate=%s\n",
		val, time.Now().Format(time.RFC3339))

	// WriteFile перезаписывает файл целиком
	if err := os.WriteFile("counter.txt", []byte(content), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "[main] ошибка записи:", err)
		os.Exit(1)
	}

	fmt.Printf("[main] сохранено: counter=%d\n", val)
}
