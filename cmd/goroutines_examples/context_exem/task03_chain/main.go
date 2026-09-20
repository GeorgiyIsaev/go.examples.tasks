package main

import (
	"context"
	"fmt"
	"time"
)

//Пример 3. Цепочка горутин и «вложенный» контекст
//Контексты наследуются: дочерний отменяется,
//если отменён родитель (или сам по таймауту).

//Оба воркера сами завершатся по одному и тому же сигналу
//— не нужно руками закрывать каналы и следить за каждым.

func producer(ctx context.Context, out chan<- int) {
	defer close(out)
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("producer: остановлен:", ctx.Err())
			return
		case out <- i:
			i++
			fmt.Println("producer: никрементирует", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func consumer(ctx context.Context, in <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("consumer: остановлен:", ctx.Err())
			return
		case v, ok := <-in:
			if !ok {
				fmt.Println("consumer: канал закрыт")
				return
			}
			fmt.Println("consumer получил:", v)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	ch := make(chan int)
	go producer(ctx, ch)
	go consumer(ctx, ch)

	<-ctx.Done() // ждём, пока сработает таймаут
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: выход, причина:", ctx.Err())
}
