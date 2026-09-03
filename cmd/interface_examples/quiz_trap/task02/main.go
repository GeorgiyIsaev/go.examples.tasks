package main

type Incr interface {
	Increment()
}

type Counter struct {
	val int
}

func (c *Counter) Increment() {
	c.val++
}

func main() {
	var i Incr
	c := Counter{} // значение
	i = c          // ошибка компиляции
	// ресивер с указтелем, передача в интерфейс как адрес
	i = &c
}
