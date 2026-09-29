package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 3 // Максимальное количество одновременных горутин

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent) // Наш семафор

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()

			// Acquire: занимаем слот
			sem <- struct{}{}
			// Release: гарантированно освобождаем слот при выходе
			defer func() { <-sem }()

			fmt.Printf("Задача %d начала выполнение\n", taskID)
			time.Sleep(1 * time.Second) // Имитация работы
			fmt.Printf("Задача %d завершена\n", taskID)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все задачи выполнены")
}
