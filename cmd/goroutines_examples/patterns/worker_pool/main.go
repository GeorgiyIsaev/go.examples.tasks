package main

import (
	"fmt"
	"sync"
	"time"
)

// worker — один воркер.
// jobs   — только читаем задачи.
// results — только пишем результаты.
// wg     — чтобы main мог дождаться завершения всех воркеров.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // воркер завершился — уменьшаем счётчик WaitGroup

	for job := range jobs { // читаем задачи, пока канал jobs не закрыт
		time.Sleep(500 * time.Millisecond) // имитация полезной работы

		results <- job * job // отправляем результат
		fmt.Printf("worker %d: обработал job %d\n", id, job)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int)    // очередь задач
	results := make(chan int) // очередь результатов

	var wg sync.WaitGroup

	// 1. Запускаем пул воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // важно сделать ДО запуска горутины
		go worker(i, jobs, results, &wg)
	}

	// 2. Отправляем задачи в отдельной горутине
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs) // сигнал: задач больше не будет
	}()

	// 3. Ждём завершения всех воркеров и только потом закрываем results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Собираем результаты, пока канал results не закрыт
	for res := range results {
		fmt.Println("result:", res)
	}
}
