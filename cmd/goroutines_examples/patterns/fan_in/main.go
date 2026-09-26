package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Создаём несколько источников (в проде — это разные БД, API, партиции Kafka)
	sources := []<-chan int{
		producer(ctx, "src-A", 1, 2, 3),
		producer(ctx, "src-B", 10, 20, 30),
		producer(ctx, "src-C", 100, 200, 300),
	}

	// 2. Fan-In: сливаем все каналы в один
	merged := fanIn(ctx, sources...)

	// 3. Потребитель читает единый поток
	for v := range merged {
		fmt.Println("получено:", v)
	}

	fmt.Println("все источники исчерпаны")
}

// fanIn сливает N входных каналов в один выходной.
func fanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// На каждый входной канал — своя горутина-переливатель
	for i, ch := range channels {
		wg.Add(1)
		go func(id int, c <-chan int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("переливатель %d: отмена\n", id)
					return
				case v, ok := <-c:
					if !ok {
						fmt.Printf("переливатель %d: источник закрыт\n", id)
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}
		}(i, ch)
	}

	// Отдельная горутина: закрыть out, когда ВСЕ переливатели закончат
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// producer имитирует источник данных.
func producer(ctx context.Context, name string, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
				time.Sleep(100 * time.Millisecond) // имитация работы
			}
		}
	}()
	return out
}
