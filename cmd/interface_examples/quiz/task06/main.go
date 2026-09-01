package main

import "fmt"

type Processor interface {
	Process(int) int
}

type Multiplier struct {
	factor int
}

func (m Multiplier) Process(x int) int {
	return x * m.factor
}

type App struct {
	p Processor
}

func (a App) Run(x int) int {
	return a.p.Process(x)
}

func main() {
	m := Multiplier{factor: 3}
	app := App{p: m}
	fmt.Println(app.Run(5)) //3*5 = 15
}
