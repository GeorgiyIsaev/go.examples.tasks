package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//Таймаут для функции через context/select
//   Задача: реализовать generic-функцию:
//func DoWithTimeout[T any](ctx context.Context, timeout time.Duration, f func() (T, error)) (T, error)
//Она запускает f в горутине.
//Если f успела — возвращает её результат.
//Если истёк таймаут или отменился ctx — возвращает ctx.Err().
//Горутина не должна блокироваться навсегда.

//Тест:
//f быстрая → результат f.
//f дольше таймаута → context.DeadlineExceeded.
//после таймаута горутина не висит, потому что канал буферизован.

func DoWithTimeout[T any](ctx context.Context, timeout time.Duration, f func() (T, error)) (T, error) {
	type result struct {
		value T
		err   error
	}

	ch := make(chan result, 1)

	go func() {
		v, err := f()
		ch <- result{v, err}
	}()

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case r := <-ch:
		return r.value, r.err
	case <-ctx.Done():
		var zero T
		return zero, ctx.Err()
	}
}

func main() {

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := DoWithTimeout(ctx, timeout, func() (int, error) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNum := r.Intn(100)
		time.Sleep(time.Duration(randomNum) * time.Second)
		return randomNum, nil
	})

	fmt.Println(result)
	fmt.Println(err)
}
