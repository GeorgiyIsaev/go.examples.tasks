package main

import "fmt"

func update(m map[string]int) {
	m["a"] = 100                 //перезапишем значение а в обеих мап
	m = map[string]int{"b": 200} // присвоение новой map
	//новая мапа как струтура будет только в функции
	//в изначальной мапе указатель останется прошлый

}

func main() {
	m := map[string]int{"a": 1, "b": 2}
	update(m)      //изменит только первое число
	fmt.Println(m) //а:100 b:2
}
