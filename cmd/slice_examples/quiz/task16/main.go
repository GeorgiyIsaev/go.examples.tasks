package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:4]         // [2,3,4], len=3, cap=4
	c := b[:2]          // [2,3], len=2, cap=4
	d := append(c, 100) // [2,3,100], len=3, cap=4
	d[0] = 200          // [200,3,100], len=3, cap=4
	a[3] = 300          // [1,200,3,300, 5], len=5, cap=5
	fmt.Println(a)      // [1,200,3,300, 5], len=5, cap=5
	fmt.Println(b)      // [200,3,300], len=3, cap=4
	fmt.Println(c)      // [200,3], len=2, cap=4
	fmt.Println(d)      // [200,3,300], len=3, cap=4
}
