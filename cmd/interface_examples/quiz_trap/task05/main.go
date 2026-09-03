package main

import "fmt"

type Changer interface {
	Change(s string)
}

type Widget struct {
	Name string
}

func (w Widget) Change(s string) {
	w.Name = s // изменения в копии не затронут оригинал
}

func main() {
	var c Changer = Widget{Name: "old"}
	c.Change("new")
	fmt.Println(c.(Widget).Name) // "old"
}
