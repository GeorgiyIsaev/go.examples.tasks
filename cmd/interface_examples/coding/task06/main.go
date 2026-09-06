package main

import (
	"bytes"
	"fmt"
	"io"
)

//Задача . Встраивание интерфейсов (композиция)
//Условие:
//Определите интерфейс ReadWriter, который встраивает в себя интерфейсы
//io.Reader и io.Writer (из пакета io). Создайте структуру MyReadWriter,
//которая содержит поле *bytes.Buffer (или встраивает bytes.Buffer).
//Реализуйте для MyReadWriter метод Read и Write (можно делегировать буферу).
//Напишите функцию Process(rw ReadWriter), которая записывает строку "Hello" в rw,
//затем читает все данные и выводит их в stdout. Продемонстрируйте работу.

type ReadWriter interface {
	io.Reader
	io.Writer
}

type MyReadWriter struct {
	*bytes.Buffer
}

func Process(rw ReadWriter) {
	_, _ = rw.Write([]byte("Hello"))
	data := make([]byte, 100)
	n, _ := rw.Read(data)
	fmt.Println(string(data[:n])) // "Hello"
}

func main() {
	buf := &bytes.Buffer{}
	mrw := MyReadWriter{Buffer: buf}
	Process(mrw)
}
