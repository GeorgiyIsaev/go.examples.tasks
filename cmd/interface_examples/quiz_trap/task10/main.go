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

func (f File) Read() string   { return "read" }
func (f File) Write(s string) { fmt.Println("write", s) }

// Допустим, у нас есть ещё интерфейс с тем же методом Read, но с другим именем?
// Встроим два интерфейса с одинаковым методом, но разными сигнатурами - ошибка компиляции.
// Но если сигнатуры одинаковы, то всё ок. Однако если они разные, то конфликт.
// Рассмотрим вариант, когда встраиваются два интерфейса с одинаковым именем метода, но разными параметрами - это ошибка.
// Например:
type Reader2 interface {
	Read(int) string
}
type ReadWriter2 interface {
	Reader  // Read() string
	Reader2 // Read(int) string
}

// Это вызовет ошибку: duplicate method Read
