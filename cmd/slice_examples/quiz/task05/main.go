package main

import "fmt"

// Вопрос: Какие значения len и cap будут выведены?
func main() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(len(s), cap(s)) //len:4 cap:6
}
