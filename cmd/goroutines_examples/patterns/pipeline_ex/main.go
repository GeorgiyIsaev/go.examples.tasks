package main

//Pipeline — это способ разложить задачу на этапы
//и выполнять их параллельно, соединив каналами.
import (
	"context"
	"fmt"
	"time"
)

// Стадия 1: генератор — источник данных
func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // обязательно закрываем, иначе следующий этап зависнет
		time.Sleep(5 * time.Second)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Println("Завершена горутина  1")
	}()
	fmt.Println("Завершена стадия 1")
	return out
}

// Стадия 2: возведение в квадрат
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		time.Sleep(5 * time.Second)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Println("Завершена горутина 2")
	}()
	fmt.Println("Завершена стадия 2")
	return out
}

// Стадия 3: фильтр — оставляем только чётные
func filterEven(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		time.Sleep(5 * time.Second)
		for n := range in {
			if n%2 != 0 {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Println("Завершена горутина 3")
	}()
	fmt.Println("Завершена стадия 3")
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Собираем pipeline: gen -> square -> filterEven
	numbers := gen(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squares := square(ctx, numbers)
	evens := filterEven(ctx, squares)

	// Потребитель (sink)
	for v := range evens {
		fmt.Println(v)
	}
	// Вывод: 4, 16, 36, 64, 100
}
