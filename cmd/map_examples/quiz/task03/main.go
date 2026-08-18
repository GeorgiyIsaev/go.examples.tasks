package main

import "fmt"

func main() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	for k, v := range m {
		fmt.Printf("%d:%s ", k, v)
	}
	//Вывод ключ значение с негарантированным порядком
}
