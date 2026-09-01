package main

import (
	"fmt"
)

// Кастомная ошибка
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("код %d: %s", e.Code, e.Msg)
}

// Функция, которая может вернуть ошибку
func doSomething(flag bool) error {
	if flag {
		return MyError{Code: 404, Msg: "не найдено"}
	}
	return nil
}

func main() {
	err := doSomething(true)
	if err != nil {
		// Утверждение типа для получения конкретной ошибки
		if me, ok := err.(MyError); ok {
			fmt.Printf("Обработана кастомная ошибка: код=%d, msg=%s\n", me.Code, me.Msg)
		} else {
			fmt.Println("Обычная ошибка:", err)
		}
	} else {
		fmt.Println("Успешно")
	}
}
