package main

import "fmt"

type Greeter interface {
	Greet()
}

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println(u.Name)
}

func main() {
	user := User{Name: "Alice"}
	//var g Greeter = user в интерфейс копируется копия структуры User.
	var g Greeter = user
	//  (var g Greeter = &user) //Если передаем указатель тогда изменения отобразатся
	user.Name = "Bob"

	g.Greet() // Alice так как внутри интерфейса копия

	/*Передаем указатель*/
	userPointer := User{Name: "Alice"}
	var gP Greeter = &userPointer
	userPointer.Name = "Bob"

	gP.Greet() // Bob так как внутри интерфейса указатель
}
