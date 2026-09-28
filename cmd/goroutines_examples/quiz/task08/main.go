package main

import (
	"fmt"
	"sync"
)

// Что выведет?
func work() {
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
}

func work_fix() {
	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
}

func main() {
	fmt.Println("Программа 1:")
	work()

	fmt.Println("Программа 2:")
	work_fix()
}
