package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (f File) Read() string {
	return "данные из файла"
}

func (f File) Write(s string) {
	fmt.Println("Запись:", s)
}

func main() {
	var rw ReadWriter = File{}
	rw.Write("привет")
	fmt.Println(rw.Read())
}
