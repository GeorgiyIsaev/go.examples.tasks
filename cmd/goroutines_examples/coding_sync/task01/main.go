package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Задача: реализовать тип Counter с методами Inc(), Add(delta int64),
//Value() int64. Мьютексы использовать нельзя.
//Тест: 100 горутин по 1000 Inc(), ожидаем 100000, запуск с -race.

type Counter struct {
	current atomic.Int64
}

func (c *Counter) Inc() {
	c.current.Add(1)
}
func (c *Counter) Add(delta int64) {
	c.current.Add(delta)
}

func (c *Counter) Value() int64 {
	return c.current.Load()
}

func main() {
	c := new(Counter)

	fmt.Println("Начало счетчика", c.Value())
	c.Inc()
	fmt.Println("Счетчик.Inc()", c.Value())

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Резульат после 100 горутин", c.Value())

}
