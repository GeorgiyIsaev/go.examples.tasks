package main

import "fmt"

type Error interface {
	Error() string
}

type MyError struct{}

func (e *MyError) Error() string { return "oops" }

func getError() Error {
	var e *MyError = nil
	return e
}

func main() {
	err := getError()
	if err != nil {
		fmt.Println("ошибка:", err) // выведет "ошибка: <nil>"
		//ошибка: oops
	} else {
		fmt.Println("нет ошибки")
	}
}
