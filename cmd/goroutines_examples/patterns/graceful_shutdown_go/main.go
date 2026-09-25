package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker инкрементирует свой локальный счётчик,
// а при отмене контекста сохраняет состояние в свой файл.
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	var counter int64   // всего инкрементов
	var iteration int64 // то же самое, но семантически "номер итерации"

	for {
		select {
		case <-ctx.Done():
			filename := fmt.Sprintf("worker_%d.txt", id)
			content := fmt.Sprintf(
				"goroutine=%d\ncounter=%d\niteration=%d\ndate=%s\n",
				id, counter, iteration, time.Now().Format(time.RFC3339),
			)
			if err := os.WriteFile(filename, []byte(content), 0o644); err != nil {
				fmt.Fprintf(os.Stderr, "[worker %d] ошибка записи %s: %v\n", id, filename, err)
				return
			}
			fmt.Printf("[worker %d] сохранено: %s (counter=%d, iter=%d)\n",
				id, filename, counter, iteration)
			return
		default:
			counter++
			iteration++
		}
	}
}

// Запускаем одноврменно 4 горутины
func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	const n = 4
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	<-ctx.Done()
	fmt.Println("\n[main] получен сигнал, дожидаемся всех горутин...")
	wg.Wait()
	fmt.Println("[main] все горутины сохранили состояние, выходим")
}
