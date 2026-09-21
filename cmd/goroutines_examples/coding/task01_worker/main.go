package main

import (
	"context"
	"fmt"
	"sync"
)

//Задача 1: Worker pool с context
//Условие: реализовать функцию, которая суммирует числа
//в N воркерах и умеет отменяться через context.Context.

func Process(ctx context.Context, nums []int, workers int) (int, error) {
	//Защищаюсь от нуля/отрицательного числа воркеров.
	//Иначе будет deadlock — никто не читает jobs
	if workers <= 0 {
		workers = 1
	}
	//Каналы
	jobs := make(chan int)
	results := make(chan int)

	//Запуск воркеров
	var wg sync.WaitGroup
	wg.Add(workers) //до запуска горутин

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
					//внешний, что бы воркер не завис при отмене контекста
				case n, ok := <-jobs:
					if !ok {
						return
					}
					select {
					case results <- n:
					case <-ctx.Done():
						return
						//внутрений, что бы при записи в results воркер
						//не завис при при отмене контекста
					}
				}
			}
		}()
	}

	//Продюсер
	go func() {
		defer close(jobs) //закрывает отправитель
		for _, n := range nums {
			select {
			case jobs <- n:
			case <-ctx.Done():
				return //что бы не завис если все воркеры вышли
			}
		}
	}()

	//Закрытие results
	go func() {
		wg.Wait() //когда все ворекры завершатся
		close(results)
	}()

	//Подсчет суммы после завершения
	sum := 0
	for v := range results {
		sum += v
	}

	if err := ctx.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	workers := 2

	ctx := context.Background()
	sum, err := Process(ctx, nums, workers)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	println("Результат", sum)

}
