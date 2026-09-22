package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Вопрос: всегда ли выведется 1000?
func ex1() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
			//!!! data race. counter++ — это не атомарная операция
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func main() {
	ex1() //!!! data race.
	ex2_mutex()
	ex3_atomic()
}

// Исправление через мьютекс
func ex2_mutex() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

// Исправление через атомик
func ex3_atomic() {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println(counter.Load())
}
