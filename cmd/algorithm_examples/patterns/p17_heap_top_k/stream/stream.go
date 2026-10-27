package main

import "container/heap"

// Потоковая статистика на кучах: два классических приёма.
//
// 1. MedianFinder (LeetCode 295): «медиана потока» на ДВУХ кучах —
//    max-heap нижней половины и min-heap верхней. Инвариант: размеры
//    различаются максимум на 1, все элементы левой <= всех правой.
//    Тогда медиана — вершина большей кучи (или среднее двух вершин).
//
// 2. MovingAverage (LeetCode 346): скользящее среднее за окно size —
//    тут куча не нужна: кольцевой буфер/срез и бегущая сумма за O(1).
//    Полезно как контраст: не каждая потоковая задача — про кучу.

// ===== Max-heap по умолчанию наоборот: нужна max-heap =====

type intMaxHeap []int

func (h intMaxHeap) Len() int           { return len(h) }
func (h intMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // больше — выше
func (h intMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intMaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder — медиана потока (LeetCode 295).
type MedianFinder struct {
	lo intMaxHeap // нижняя половина (max-heap: её максимум — кандидат)
	hi intMinHeap // верхняя половина (min-heap: её минимум — кандидат)
}

func NewMedianFinder() *MedianFinder {
	m := &MedianFinder{}
	heap.Init(&m.lo)
	heap.Init(&m.hi)
	return m
}

// AddNum добавляет число в поток, сохраняя инвариант куч:
// lo (нижняя половина, max-heap) не меньше hi, разница в размерах <= 1,
// lo[0] <= hi[0].
func (m *MedianFinder) AddNum(num int) {
	heap.Push(&m.lo, num) // 1. сначала в нижнюю половину
	// 2. если нарушили порядок — верхняя граница ниже нижней, размен.
	if m.hi.Len() > 0 && m.lo[0] > m.hi[0] {
		heap.Push(&m.hi, heap.Pop(&m.lo).(int))
	}
	// 3. балансируем размеры: lo больше hi не больше чем на 1.
	if m.lo.Len() > m.hi.Len()+1 {
		heap.Push(&m.hi, heap.Pop(&m.lo).(int))
	} else if m.hi.Len() > m.lo.Len() {
		heap.Push(&m.lo, heap.Pop(&m.hi).(int))
	}
}

// FindMedian возвращает текущую медиану.
func (m *MedianFinder) FindMedian() float64 {
	if m.lo.Len() > m.hi.Len() {
		return float64(m.lo[0])
	}
	return float64(m.lo[0]+m.hi[0]) / 2
}

// intMinHeap — min-heap верхней половины (пакет отдельный от top_k,
// поэтому определение дублируется).
type intMinHeap []int

func (h intMinHeap) Len() int           { return len(h) }
func (h intMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intMinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *intMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ===== Moving Average (контраст: без кучи) =====

// MovingAverage — скользящее среднее за окно size (LeetCode 346).
type MovingAverage struct {
	size   int
	window []int
	sum    int
}

func NewMovingAverage(size int) *MovingAverage {
	return &MovingAverage{size: size}
}

// Next возвращает среднее по последним size значениям.
func (m *MovingAverage) Next(val int) float64 {
	m.window = append(m.window, val)
	m.sum += val
	if len(m.window) > m.size {
		m.sum -= m.window[0]
		m.window = m.window[1:]
	}
	return float64(m.sum) / float64(len(m.window))
}
