package main

import (
	"fmt"
	"sync"
	"time"
)

// Что не так с использованием WaitGroup
func ex() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done() // counter = 0, Wait может проснуться
		wg.Add(1) // новая партия
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Println("second done")
		}()
	}()
	//Нарушен порядок
	wg.Wait()
	fmt.Println("main done")
}

// Правильно:
func ex_true() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Println("second done")
		}()
	}()

	wg.Wait()
	fmt.Println("main done")
}

func main() {
	fmt.Println("Запуск 1")
	//ex() // тут ошибка
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Запуск 2")
	ex_true()
	time.Sleep(500 * time.Millisecond)
}
