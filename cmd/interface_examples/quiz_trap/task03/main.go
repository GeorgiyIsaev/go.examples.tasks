package main

import "fmt"

type Printer interface {
	Print()
}

type Data struct {
	msg string
}

func (d Data) Print() {
	fmt.Println(d.msg)
}

func main() {
	var p Printer
	d := &Data{msg: "hello"}
	p = d     //автоматическое разименование - в метод отправлена копия значения
	p.Print() // выведет "hello"
}
