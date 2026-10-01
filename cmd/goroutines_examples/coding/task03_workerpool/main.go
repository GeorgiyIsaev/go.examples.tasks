package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Задача 3: параллельная обработка URL с ограничением
//Условие. Дан список URL (например, 1000 штук). Нужно:
//- Скачать каждый URL (заглушка — fetch(url), возвращает len контента)
//- Ограничить параллелизм — не больше N одновременных запросов (например, 10)
//- Собрать результаты и посчитать суммарный размер скачанного
//- Если какой-то URL упал — не ронять всю задачу, а собрать ошибки
//- Поддержать отмену через context

//ПЛАН:

//Заглушка для — fetch(url) - будет имитировать сетевую работу и ошибки
//func fetch(ctx context.Context, url string) (int, error)

// processURLs(контекст, юрл стринг, количество ворекров) (результат, ошибка)
//- метод для ворекрпула с ограничением
//джобс и резеалт каналы и запуск горутин

//В маин:
// Создадим слайс юрл стринг на 50 штук
// Контест по времени на 3 секунды
// запустим воркер processURLs
// выведем результат
// проверим ошибки и прерывания

//РЕШЕНИЕ:

// Результат обработки одного юрл
type Result struct {
	URL  string
	Size int
	Err  error
}

// Заглушка вместо реально HHTP-зпроса
func fetch(ctx context.Context, url string) (int, error) {
	//имитация задержки
	select {
	case <-time.After(time.Duration(50+rand.Intn(150)) * time.Millisecond):
	case <-ctx.Done():
		return 0, ctx.Err()
	}
	//имитация случайно ошибки
	if rand.Intn(10) == 0 {
		return rand.Intn(100), fmt.Errorf("failed to fetch: %s", url)
	}
	return 100 + rand.Intn(900), nil
}

// processURLs — воркерпул с ограничением параллелизма N
func processURLs(ctx context.Context, urls []string, workers int) (sum int, err []error) {
	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup

	//1. ЗАПУСКАМ N ворекров
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
				}

				size, err := fetch(ctx, j)
				results <- Result{j, size, err}
			}
		}()
	}

	//2. Запускаем работы в отельной горутине
	go func() {
		defer close(jobs)
		for _, url := range urls {
			select {
			case <-ctx.Done():
				return
			case jobs <- url:
			}
		}
	}()

	//3 Горутина для закрытия результатот когда все ворекры закончили работу
	go func() {
		wg.Wait()
		close(results)
	}()

	//4 Сбор результата
	var (
		totalSize int
		errs      []error
	)
	for r := range results {
		if r.Err != nil {
			errs = append(errs, r.Err)
		}
		totalSize += r.Size
	}
	return totalSize, errs
}

func main() {
	urls := make([]string, 50)
	for i := range urls {
		urls[i] = fmt.Sprintf("http://exemlle.com/%d", i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	total, errs := processURLs(ctx, urls, 10)

	fmt.Println("Суммарный размер:", total)
	fmt.Println("Ошибок:", len(errs))

	for _, e := range errs {
		fmt.Println(" -", e)
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		fmt.Println("Часть работы была прервана по таймауту")
	}
}
