package main

import "fmt"

func cange(s []int) {
	s[0] = 100
}
func main() {
	a := []int{1, 2, 3}
	cange(a)
	fmt.Println(a) //100,2,3
}
