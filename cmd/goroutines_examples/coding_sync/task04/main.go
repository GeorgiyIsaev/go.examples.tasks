package main

import (
	"context"
	"fmt"
	"sync"
)

//Worker pool с jobs/results и context
//Задача: реализовать функцию:
//func WorkerPool(ctx context.Context, jobs <-chan int, workers int, f func(int) int) <-chan int
//- Она запускает workers воркеров, читает jobs,
//пишет результаты в выходной канал и закрывает его
//после завершения всех воркеров. Должна уважать отмену ctx.
//- Тест: закрыть jobs, прочитать всё из out,
//проверить отсутствие паники и утечек.

func WorkerPool(ctx context.Context, jobs <-chan int, workers int, f func(int) int) <-chan int {
	results := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}

					select {
					case results <- f(job):
					case <-ctx.Done():
						return
					}
				}

			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs := make(chan int, 100)
	results := WorkerPool(ctx, jobs, 10, func(i int) int { return i * i })

	go func() {
		defer close(jobs)
		for i := 0; i < 100; i++ {
			jobs <- i
		}
	}()

	for result := range results {
		fmt.Println(result)
	}

}
