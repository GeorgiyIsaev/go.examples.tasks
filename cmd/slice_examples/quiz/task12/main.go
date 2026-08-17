package main

import "fmt"

func main() {
	s := make([]int, 0, 5)                 //len0 сap5
	s = append(s, 1, 2, 3)                 //1,2,3 len3 сap5
	t := s[:2]                             //1,2 len2 сap5
	t = append(t, 100, 200, 300, 400, 500) //1,2, 100,200,300,400,500 len7 сap10 (новый)
	fmt.Println(s)                         //1,2,3 len3 сap5
	fmt.Println(t)                         //1,2, 100,200,300,400,500 len7 сap10
}
