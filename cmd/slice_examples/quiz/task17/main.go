package main

import "fmt"

func modify(s []int) []int {
	s = append(s, 4) // [1,2,3,4], len=4, cap=5
	s[0] = 100       // [100,2,3,4], len=4, cap=5
	return s
}

func main() {
	a := make([]int, 3, 5)      // [], len=3, cap=5
	a[0], a[1], a[2] = 1, 2, 3  // [1,2,3], len=3, cap=5
	b := modify(a)              // [100,2,3,4], len=4, cap=5 (и а и b на один участок)
	a[1] = 200                  // [100,200,3], len=3, cap=5
	b = append(b, 5)            // [100,200,3,4,5], len=5, cap=5
	b[2] = 300                  // [100,200,300,4,5], len=5, cap=5
	fmt.Println(a)              // [100,200,300] len=5, cap=5
	fmt.Println(b)              // [100,200,300,4,5], len=5, cap=5
	fmt.Println(cap(a), cap(b)) // a cap=5 b cap=5
}
