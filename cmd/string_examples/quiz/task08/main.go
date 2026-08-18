package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "hello"
	//unsafe.StringData(s) возвращает указатель на первый байт
	//внутреннего массива строки, который обычно находится в read-only сегменте.
	//unsafe.Slice создаёт срез, указывающий на эти данные, но их нельзя менять
	b := unsafe.Slice(unsafe.StringData(s), len(s))
	b[0] = 'H' // попытка изменить вызывает неопределенное поведение
	fmt.Println(s)
	//произойдет паника
}
