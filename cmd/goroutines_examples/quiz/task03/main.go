package main

import "fmt"

// Что выведет?
func test1() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	//deadlock. Отправка в небуферизованный канал блокируется.
}

func main() {
	//test1()	//deadlock
	test2()
	test3()
}

// Исправление: добавляем буффер
func test2() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch) //1
}

// Исправление: Или можно добавить занчение в канал в другой горутине
// будет заблакирована внутреняя горутна
func test3() {
	ch := make(chan int)
	go func() { ch <- 1 }()
	fmt.Println(<-ch) //1
}
