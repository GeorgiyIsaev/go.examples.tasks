package main

// Паттерн: куча (heap) для Top-K задач.
//
// Идея: вместо полной сортировки O(n log n) держим в куче только k
// «кандидатов» и гоняем по n элементов через heap.Push/Pop:
// O(n log k) — выигрыш, когда k << n.
//
// Какой кучей:
//   - ищем k-й САМЫЙ БОЛЬШОЙ -> **min-heap из k элементов**:
//     вершина = граница топ-k, новый кандидат сравнивается с ней;
//   - нужен k-й МЕНЬШИЙ -> max-heap из k;
//   - поток и «среднее/топ по частоте» — комбинируем.
//
// В Go — container/heap: реализуешь heap.Interface (Len/Less/Swap/
// Push/Pop) и вызываешь heap.Init/heap.Push/heap.Pop.
// Less(i,j)=true значит «i вверху»: для min-heap h[i] < h[j],
// для max-heap — наоборот.
//
// Сложность: O(n log k) на топ-K, память O(k).

import (
	"container/heap"
)

// ===== min-heap из целых =====

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

// KthLargest — k-й наибольший элемент массива (LeetCode 215).
// k считается среди элементов с дублями (1-based).
//
// Алгоритм: min-heap ровно из k элементов. Вершина = k-й большой
// из виденных. Новый элемент больше вершины? вытесняет её.
// В конце вершина и есть ответ. O(n log k).
//
// Пример: [3,2,1,5,6,4], k=2 -> 5.
func KthLargest(nums []int, k int) int {
	h := &intMinHeap{}
	heap.Init(h)
	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
			continue
		}
		if v > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	return (*h)[0]
}

// KthLargestStream — «k-й наибольший в потоке» (LeetCode 703):
// конструктор с k, метод Add за O(log k).
//
// Пример: KthL2(703-шный) new(3), Add(10)=10, Add(5)=7, Add(4)=7...
type KthLargestStream struct {
	k int
	h intMinHeap
}

func NewKthLargestStream(k int, nums []int) *KthLargestStream {
	s := &KthLargestStream{k: k}
	heap.Init(&s.h)
	for _, v := range nums {
		s.Add(v)
	}
	return s
}

// Add возвращает k-й наибольший после добавления val.
func (s *KthLargestStream) Add(val int) int {
	if s.h.Len() < s.k {
		heap.Push(&s.h, val)
	} else if val > s.h[0] {
		heap.Pop(&s.h) // вытесняем худшего из топ-k
		heap.Push(&s.h, val)
	}
	if s.h.Len() == 0 {
		return 0 // k <= 0 по условию не бывает, для безопасности
	}
	return s.h[0]
}

// ===== куча по произвольному ключу =====

type pointItem struct {
	idx  int // индекс в исходном срезе
	dist int // квадрат расстояния до начала координат
}

type pointMaxHeap []pointItem

func (h pointMaxHeap) Len() int           { return len(h) }
func (h pointMaxHeap) Less(i, j int) bool { return h[i].dist > h[j].dist } // max по dist
func (h pointMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pointMaxHeap) Push(x any)        { *h = append(*h, x.(pointItem)) }
func (h *pointMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KClosest — k точек, ближайших к началу координат (LeetCode 973).
// Порядок ответа не важен; ключ кучи — квадрат расстояния
// (корень не нужен, сравнения те же).
//
// Два варианта: max-heap из k (как здесь) или min-heap из всех n
// и снять k вершин. При k << n — первый выгоднее.
//
// Пример: [[1,3],[-2,2]], k=1 -> [[-2,2]].
func KClosest(points [][]int, k int) [][]int {
	h := &pointMaxHeap{}
	heap.Init(h)
	for i, p := range points {
		d := p[0]*p[0] + p[1]*p[1]
		if h.Len() < k {
			heap.Push(h, pointItem{i, d})
			continue
		}
		if d < (*h)[0].dist { // ближе худшего из топ-k — вытесняем
			heap.Pop(h)
			heap.Push(h, pointItem{i, d})
		}
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 {
		it := heap.Pop(h).(pointItem)
		result = append(result, points[it.idx])
	}
	return result
}
