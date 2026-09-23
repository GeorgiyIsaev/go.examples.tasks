package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//написать функцию, которая ждёт значение из канала,
//но не дольше, чем позволяет контекст.

func Wait(ctx context.Context, ch <-chan int) (int, error) {
	select {
	case v, ok := <-ch:
		if !ok {
			return 0, errors.New("channel closed")
		}
		return v, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	scenario_onTime()
	scenario_timeout()
	scenario_cancel()

}
func scenario_onTime() {
	// --- Сценарий 1: значение пришло вовремя ---
	ch := make(chan int, 1) // буферизованный, чтобы отправитель не блокировался
	ch <- 42                // кладём значение сразу
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	v, err := Wait(ctx, ch)
	if err != nil {
		fmt.Println("сценарий 1: ошибка:", err)
	} else {
		fmt.Println("сценарий 1: получили значение:", v)
	}
}

func scenario_timeout() {
	// --- Сценарий 2: значение не пришло, сработал таймаут ---
	ch := make(chan int) // пустой, никто не пишет

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	start := time.Now()
	v, err := Wait(ctx, ch)
	fmt.Printf("сценарий 2: ждали %v, v=%d, err=%v\n",
		time.Since(start).Round(time.Millisecond), v, err)
}

func scenario_cancel() {
	// --- Сценарий 3: канал закрыт ---
	ch := make(chan int)
	close(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	v, err := Wait(ctx, ch)
	fmt.Println("сценарий 3: v=", v, "err=", err)
}
