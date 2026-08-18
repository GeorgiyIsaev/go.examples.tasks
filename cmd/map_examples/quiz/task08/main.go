package main

import "fmt"

func main() {
	m := map[int]bool{1: true, 2: true}
	for k := range m {
		if k == 1 {
			m[3] = true // добавляем новый ключ

		}
	}
	//паники не будет,
	//но результат не гарантирован
	fmt.Println(m) // вероятно map[1:true 2:true 3:true] а может и 1,2
}
