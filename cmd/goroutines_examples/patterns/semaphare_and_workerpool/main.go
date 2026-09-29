package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	totalTasks    = 50
	workerCount   = 10
	maxConcurrent = 3
	taskDuration  = 300 * time.Millisecond
)

// worker обрабатывает задачи из канала jobs, ограничивая параллелизм через sem.
func worker(
	workerID int,
	jobs <-chan int,
	sem chan struct{},
	active *int32,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for taskID := range jobs {
		sem <- struct{}{} // захват семафора

		nowActive := atomic.AddInt32(active, 1)
		printStart(workerID, taskID, nowActive)

		doWork()

		atomic.AddInt32(active, -1)
		<-sem // освобождение семафора

		printDone(workerID, taskID)
	}
}

// doWork имитирует полезную нагрузку.
func doWork() {
	time.Sleep(taskDuration)
}

// printStart печатает сообщение о начале обработки задачи.
func printStart(workerID, taskID int, active int32) {
	fmt.Printf("worker %2d | task %2d | START (active=%d)\n",
		workerID, taskID, active)
}

// printDone печатает сообщение о завершении обработки задачи.
func printDone(workerID, taskID int) {
	fmt.Printf("worker %2d | task %2d | DONE\n", workerID, taskID)
}

// startWorkers запускает пул воркеров и возвращает WaitGroup для ожидания.
func startWorkers(
	jobs <-chan int,
	sem chan struct{},
	active *int32,
) *sync.WaitGroup {
	var wg sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, sem, active, &wg)
	}
	return &wg
}

// generateJobs наполняет канал jobs задачами и закрывает его.
func generateJobs(jobs chan<- int) {
	for i := 1; i <= totalTasks; i++ {
		jobs <- i
	}
	close(jobs)
}

func main() {
	jobs := make(chan int, totalTasks)
	sem := make(chan struct{}, maxConcurrent)
	var active int32

	wg := startWorkers(jobs, sem, &active)
	generateJobs(jobs)
	wg.Wait()

	fmt.Println("Все 50 задач обработаны")
}
