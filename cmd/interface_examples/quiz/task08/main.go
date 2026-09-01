package main

type Incrementer interface {
	Inc()
}

type Counter struct {
	val int
}

// Метод с указателем-получателем
func (c *Counter) Inc() {
	c.val++
}

func main() {
	var inc Incrementer

	c := Counter{} // значение
	inc = c        // ошибка: Counter does not implement Incrementer (method Inc has pointer receiver)
	_ = inc

	//Если метод с указателем, то интерфейс через адрес
	//inc = &c // теперь ок
	//_ = inc
}
