package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.TODO()
	maxWorkers := runtime.GOMAXPROCS(0) // Количество логических CPU

	// Создаем взвешенный семафор с общим весом maxWorkers
	sem := semaphore.NewWeighted(int64(maxWorkers))
	out := make([]int, 32)

	for i := range out {
		// Acquire блокируется, если нет свободного веса
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Не удалось захватить семафор: %v", err)
			break
		}
		go func(i int) {
			defer sem.Release(1) // Освобождаем 1 единицу веса
			out[i] = collatzSteps(i + 1)
		}(i)
	}

	// Ожидание завершения всех горутин: захватываем весь доступный вес
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Не удалось захватить семафор: %v", err)
	}
	fmt.Println(out)
}

func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}
	for ; n > 1; steps++ {
		if n%2 == 0 {
			n /= 2
			continue
		}
		n = 3*n + 1
	}
	return steps
}
