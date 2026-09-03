package main

import "fmt"

func main() {
	var x any //тоже самое что и пустой интерфейс
	var p *int = nil
	x = p
	fmt.Println(x == nil) // false
	//хотя значение nil сам интерфейс не nil
}
