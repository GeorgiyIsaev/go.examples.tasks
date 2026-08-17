package main

import "fmt"

// Что будет выведено? Опишите, как работают слайсы слайсов и вложенные append.
func main() {
	s := [][]int{
		{1, 2},
		{3, 4},
	}
	s = append(s, []int{5, 6})
	s[0] = append(s[0], 7)
	fmt.Println(s)
}
