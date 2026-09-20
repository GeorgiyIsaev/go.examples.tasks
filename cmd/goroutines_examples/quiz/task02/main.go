package main

import (
	"fmt"
	"sync"
)

// Что не так в коде?
func task_noAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}

func main() {
	//task_noAdd()
	task_withAdd()
}

// Исправление
func task_withAdd() {
	var wg sync.WaitGroup //создает ваит группу
	//Нужно добавить wg.Add(5) -- увеличивает счётчик ожидаемых задач/горутин на 5.
	wg.Add(5) //5 горутин в группе

	for i := 0; i < 5; i++ {
		go func(i int) { //i передаем как копию в горутину
			defer wg.Done() //выполняется когда горутина завершится
			fmt.Println(i)
		}(i)
	}
	//Вывод горутин в случайном порядке
	wg.Wait() //блокирует ожидание пока все горутны из группы не завершатся
}
