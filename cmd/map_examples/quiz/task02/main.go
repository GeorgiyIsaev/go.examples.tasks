package main

import "fmt"

func main() {
	m := map[string]int{"x": 10, "y": 20}
	delete(m, "x") // удалим х
	delete(m, "z") // удаление несуществующего ключа — безопасно
	fmt.Println(m["x"], m["y"])
	//0 20
}
