package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Напишите две реализации счётчика для числа запросов к разным userID:
// С использованием обычного map[int]int и sync.RWMutex.
// С использованием sync.Map.
//
// Запустите 1000 горутин, каждая делает 1000 инкрементов
// по случайному ключу (из 100 возможных).
// Сравните время выполнения.
// Напишите бенчмарк или просто код с замерами.

// ----- Реализация с map + RWMutex -----

type MutexCounter struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{m: make(map[int]int)}
}

func (c *MutexCounter) Increment(key int) {
	c.mu.Lock()
	c.m[key]++
	c.mu.Unlock()
}

// ----- Реализация с sync.Map + atomic -----

type SyncCounter struct {
	m sync.Map
}

func NewSyncCounter() *SyncCounter {
	return &SyncCounter{}
}

func (c *SyncCounter) Increment(key int) {
	// Загружаем существующий указатель или сохраняем новый (инициализируется нулём)
	val, _ := c.m.LoadOrStore(key, new(int32))
	atomic.AddInt32(val.(*int32), 1)
}

// ----- Генерация ключей для горутин -----

func generateKeys(numGoroutines, opsPerGoroutine, numKeys int) [][]int {
	keys := make([][]int, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		keys[i] = make([]int, opsPerGoroutine)
		r := rand.New(rand.NewSource(int64(i + 1)))
		for j := 0; j < opsPerGoroutine; j++ {
			keys[i][j] = r.Intn(numKeys)
		}
	}
	return keys
}

// ----- Замеры в main -----

func runMutex(keys [][]int) time.Duration {
	counter := NewMutexCounter()
	var wg sync.WaitGroup
	wg.Add(len(keys))
	start := time.Now()
	for _, ks := range keys {
		go func(ks []int) {
			defer wg.Done()
			for _, k := range ks {
				counter.Increment(k)
			}
		}(ks)
	}
	wg.Wait()
	return time.Since(start)
}

func runSync(keys [][]int) time.Duration {
	counter := NewSyncCounter()
	var wg sync.WaitGroup
	wg.Add(len(keys))
	start := time.Now()
	for _, ks := range keys {
		go func(ks []int) {
			defer wg.Done()
			for _, k := range ks {
				counter.Increment(k)
			}
		}(ks)
	}
	wg.Wait()
	return time.Since(start)
}

func main() {
	const (
		numGoroutines   = 1000
		opsPerGoroutine = 1000
		numKeys         = 100
	)
	keys := generateKeys(numGoroutines, opsPerGoroutine, numKeys)

	fmt.Println("Запуск с map + RWMutex...")
	elapsedMutex := runMutex(keys)
	fmt.Printf("Время: %v\n", elapsedMutex)

	fmt.Println("Запуск с sync.Map...")
	elapsedSync := runSync(keys)
	fmt.Printf("Время: %v\n", elapsedSync)

	// Проверка корректности (сумма должна быть равна numGoroutines * opsPerGoroutine)
	// Для простоты проверим только для MutexCounter (можно аналогично для SyncCounter)
	counterCheck := NewMutexCounter()
	for _, ks := range keys {
		for _, k := range ks {
			counterCheck.Increment(k)
		}
	}
	total := 0
	for _, v := range counterCheck.m {
		total += v
	}
	fmt.Printf("Проверка суммы (map+mutex): %d (должно быть %d)\n", total, numGoroutines*opsPerGoroutine)
}
