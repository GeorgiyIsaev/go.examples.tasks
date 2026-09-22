package main

import (
	"context"
	"fmt"
	"time"
)

// Вопрос: что не так?
func ex(ctx context.Context) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
		return
	case v := <-ch:
		fmt.Println(v)
	}
	//ТУТ Утечка горутины
	//Done сработает раньше маин
	//горутина навсегда заблокируется на ch <- 1
}

func main() {
	//Сценарий 1: Утечка горутины дон сработает раньше маин
	fmt.Println("Сценарий 1:")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go ex(ctx)
	//cancel()
	time.Sleep(500 * time.Millisecond)

	//Сценарий 2: Буфер
	fmt.Println("Сценарий 2:")
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()
	go ex_buffer(ctx2)
	//cancel2()
	time.Sleep(500 * time.Millisecond)

	//Сценарий 3: select внутри горутины
	fmt.Println("Сценарий 3:")
	ctx3, cancel3 := context.WithCancel(context.Background())
	defer cancel3()
	go ex_select(ctx3)
	//cancel3()
	time.Sleep(500 * time.Millisecond)

}

// Решение 1: Сделать буфер
func ex_buffer(ctx context.Context) {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
		return
	case v := <-ch:
		fmt.Println(v)
	}

}

// Решение 2: select внутри горутины
func ex_select(ctx context.Context) {
	ch := make(chan int)

	go func() {
		select {
		case ch <- 1:
		case <-ctx.Done():
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timed out")
		return
	case v := <-ch:
		fmt.Println(v)
	}
}
