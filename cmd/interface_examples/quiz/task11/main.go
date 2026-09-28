package main

import "fmt"

type MyError struct{}

func (*MyError) Error() string {
	return "my error"
}

func main() {
	var e *MyError
	var err error = e

	fmt.Println(e == nil)   //nil
	fmt.Println(err == nil) //не нил, так интфейс хранит данные о самой переменой
}
