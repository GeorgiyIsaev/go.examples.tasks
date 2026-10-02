package main

import (
	"fmt"
	"sync"
)

//   Задача: написать функцию ParallelMap(in []int, f func(int) int) []int,
//   которая обрабатывает каждый элемент в отдельной горутине,
//   но результат возвращает в исходном порядке.
//   Тест: ParallelMap([]int{1,2,3}, func(x int) int { return x * x }) → [1,4,9], -race.

func ParallelMap(in []int, f func(int) int) []int {
	out := make([]int, len(in))

	var wg sync.WaitGroup
	wg.Add(len(in))

	for i, v := range in {
		go func(i int, v int) {
			defer wg.Done()
			out[i] = f(v)
		}(i, v)
	}

	wg.Wait()
	return out
}

func main() {

	out := ParallelMap([]int{1, 2, 3}, func(x int) int { return x * x })
	fmt.Println(out)
}
