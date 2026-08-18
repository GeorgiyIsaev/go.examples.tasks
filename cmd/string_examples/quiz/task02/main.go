package main

import "fmt"

func main() {
	s := ""
	for i := 0; i < 5; i++ {
		s += "a" //какждый раз пересоздание
	}
	fmt.Println(s) //aaaaa (5 a)
}
