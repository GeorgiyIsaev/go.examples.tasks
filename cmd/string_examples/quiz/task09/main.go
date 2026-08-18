package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "hello"
	b := "hello"      // извлекает литерал из существующей памяти
	c := "hel" + "lo" // константная конкатенация также
	d := "hel"        // кусок из памяти
	e := d + "lo"     // переменная + константа // новая память

	fmt.Println(a == b) //true по значению
	fmt.Println(a == c) //true по значению
	fmt.Println(a == e) //true по значению

	// Сравним указатели (через unsafe)
	fmt.Println(*(*string)(unsafe.Pointer(&a)) == *(*string)(unsafe.Pointer(&b))) // так не сравнить указатели, лучше через StringHeader
	// Используем reflect.StringHeader (устаревший, но для демонстрации)
	ha := (*reflect.StringHeader)(unsafe.Pointer(&a))
	hb := (*reflect.StringHeader)(unsafe.Pointer(&b))
	hc := (*reflect.StringHeader)(unsafe.Pointer(&c))
	he := (*reflect.StringHeader)(unsafe.Pointer(&e))
	fmt.Println(ha.Data == hb.Data) //true одинаковый участок памяти
	fmt.Println(ha.Data == hc.Data) //true одинаковый участок памяти
	fmt.Println(ha.Data == he.Data) //false новый участок памяти
}
