package main

import (
	"fmt"
	"sync"
	"time"
)

// Fan-Out: Веерное слияние
func main() {
	jobs := make(chan int, 100)
	fmt.Println("МАИН: Запуск")
	// Источник задач
	go func() {
		fmt.Println("Продюсер: Запуск")
		for i := 1; i <= 20; i++ {
			jobs <- i
		}
		fmt.Println("Продюсер: Ждем 5s")
		time.Sleep(5 * time.Second)
		fmt.Println("Продюсер: Перед Клозе")
		close(jobs)
		fmt.Println("Продюсер: Конец")
	}()

	//fmt.Println("Маин: Ждем 5s что бы запустить воркер")
	//time.Sleep(15 * time.Second)
	//fmt.Println("Маин: Переходи к воркерам")

	// Fan-Out: 3 воркера читают из одного канала
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		fmt.Println("МАИН: Запуску воркера: ", w)
		wg.Add(1)
		go func(workerID int) {
			fmt.Printf("воркер %d запущен\n", workerID)
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("воркер %d обработал %d\n", workerID, job)
			}
			fmt.Printf("воркер %d завершен\n", workerID)
		}(w)
	}

	fmt.Println("МАИН: Ждем завершения")
	wg.Wait()
	fmt.Println("МАИН: Конец")
}
