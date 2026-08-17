package main

import "fmt"

// что вывдем
func main() {
	a := make([]int, 0, 3) // len:0 cap:3 пустой срез
	a = append(a, 1, 2)    // len:2 cap:3 //заполняем элементом 1 и 2

	b := append(a, 3) // slice a: l2 c3, slice b: l3 c3
	c := append(a, 4) // slice a: l2 c3, slice c: l3 c3

	fmt.Println(b)     //1 2 4  // и b и с обращаются к одному участку памяти при изменении в с, b тоже меняется
	fmt.Println(c)     //1 2 4
	fmt.Println(a[:3]) // 1 2 4 //а тоже меняется, если увеличить длину мы это увидим.
}
