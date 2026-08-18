package main

import "fmt"

func main() {
	s := "abc\x00def"   //символ окончания строки не влиет на строку у неё есть свой размер
	fmt.Println(len(s)) //все символы abc(3) \x00(1) def(3) = 7
	fmt.Println(s)      //abc0def
	fmt.Println(s[3])   //0 - символ ничто
}
