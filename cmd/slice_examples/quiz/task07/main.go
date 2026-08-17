package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	sub := s[1:3]
	fmt.Println(len(sub), cap(sub)) //len 2 cap 4
	fmt.Println(sub)                //[2,3]

	// capacity sub равна 4.
	//вычисляется как разность между ёмкостью исходного среза
	//и начальным индексом нового среза 5-1=4
	//Другой пример
	s2 := []int{10, 20, 30, 40} // len=4, cap=4
	sub2 := s2[0:2]             // len=2, cap = 4-0 = 4
	fmt.Println(cap(sub2))      // 4

	//
	s3 := []int{1, 2, 3, 4, 5}
	sub3 := s3[1:3:4]      // len=2, cap = 4-1 = 3
	fmt.Println(cap(sub3)) // 3
}
