package main

import "fmt"

func main() {
	s := make([]int, 0, 3) //len 0 cap 3
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("len=%d cap=%d\n", len(s), cap(s)) //len 5 cap 6
	}
}
