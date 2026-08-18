package main

import "fmt"

func main() {
	s := "мир"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // байтовые представления 6штук
		//перевод в с случайный символ
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c ", r) //м и р - нормальные символы
		// range декодирует UTF-8
	}
}
