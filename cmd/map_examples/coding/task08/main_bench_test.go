package main

import (
	"sync"
	"testing"
)

// Для бенчмарков используем те же параметры, но ключи генерируются один раз вне цикла.
func BenchmarkMutex(b *testing.B) {
	const (
		numGoroutines   = 1000
		opsPerGoroutine = 1000
		numKeys         = 100
	)
	keys := generateKeys(numGoroutines, opsPerGoroutine, numKeys)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := NewMutexCounter()
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		for _, ks := range keys {
			go func(ks []int) {
				defer wg.Done()
				for _, k := range ks {
					counter.Increment(k)
				}
			}(ks)
		}
		wg.Wait()
	}
	b.StopTimer()
}

func BenchmarkSync(b *testing.B) {
	const (
		numGoroutines   = 1000
		opsPerGoroutine = 1000
		numKeys         = 100
	)
	keys := generateKeys(numGoroutines, opsPerGoroutine, numKeys)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := NewSyncCounter()
		var wg sync.WaitGroup
		wg.Add(numGoroutines)
		for _, ks := range keys {
			go func(ks []int) {
				defer wg.Done()
				for _, k := range ks {
					counter.Increment(k)
				}
			}(ks)
		}
		wg.Wait()
	}
	b.StopTimer()
}
