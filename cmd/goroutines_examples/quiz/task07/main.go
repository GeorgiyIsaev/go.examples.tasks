package main

import (
	"fmt"
	"sync"
)

// Вопрос: что произойдёт?
func ex() {
	var once sync.Once

	once.Do(func() {
		fmt.Println("outer")
		once.Do(func() {
			fmt.Println("inner")
		})
	})

	fmt.Println("done")
	//Ответ: deadlock. sync.Once.Do держит внутренний mutex
}

func main() {
	//ex() // deadlock
	ex_true()
}

// Исправление:
func ex_true() {
	var once sync.Once

	init := func() {
		fmt.Println("outer")
		fmt.Println("inner")
	}

	once.Do(init) //Do вызовется один раз
	fmt.Println("done")
	//Ответ: deadlock. sync.Once.Do держит внутренний mutex
}
