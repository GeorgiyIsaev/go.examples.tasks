package main

import "fmt"

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Целое: %d\n", v) // Целое: 10
	case string:
		fmt.Printf("Строка: %s\n", v) // Строка: hello
	default:
		fmt.Printf("Другой тип: %T\n", v) // Другой тип: float64
	}
}

func main() {
	describe(10)
	describe("hello")
	describe(3.14)
}
