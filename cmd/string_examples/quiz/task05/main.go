package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s)  //копия байтов 6рун =6 символов
	b[0] = 'H'      //замена байта и в 1 руну успешна
	s2 := string(b) //перевод байтов в строку новая строка
	fmt.Println(s)  //hello
	fmt.Println(s2) //Hello
}
