package main

import "container/heap"

// Слияние k отсортированных списков — классическая задача на кучу
// (LeetCode 23 Merge k Sorted Lists).
//
// Идея: k-стороннее слияние. Вместо поиска минимума среди k списков
// за O(k) куча даёт O(log k) на каждый шаг:
//
//	1. В кучу — первые элементы всех списков (k штук).
//	2. Забрали минимум (вершина) — записали в результат.
//	3. Взяли следующий из того же списка — положили в кучу.
//	4. Повторяли, пока куча не опустеет.
//
// Сложность: O(N log k), N — всего элементов, k — списков.
// (Альтернатива — попарное слияние за O(N log k) тоже, но с бо́льшими
// константами и без кучи.)
//
// ListNode определён в list_node.go этой же папки.

// listItem — элемент кучи: значение и указатель на «свой» узел.
type listItem struct {
	val  int
	node *ListNode
}

type listMinHeap []listItem

func (h listMinHeap) Len() int           { return len(h) }
func (h listMinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h listMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *listMinHeap) Push(x any)        { *h = append(*h, x.(listItem)) }
func (h *listMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MergeKLists — сливает k отсортированных списков в один.
// Пустые списки игнорируются; если все пусты — nil.
//
// Пример: [1,4,5], [1,3,4], [2,6] -> [1,1,2,3,4,4,5,6].
func MergeKLists(lists []*ListNode) *ListNode {
	h := &listMinHeap{}
	heap.Init(h)
	for _, l := range lists {
		if l != nil {
			heap.Push(h, listItem{l.Val, l})
		}
	}

	dummy := &ListNode{} // фиктивная голова, чтобы не размножать if
	tail := dummy
	for h.Len() > 0 {
		it := heap.Pop(h).(listItem)
		tail.Next = it.node
		tail = tail.Next
		if it.node.Next != nil {
			heap.Push(h, listItem{it.node.Next.Val, it.node.Next})
		}
	}
	return dummy.Next
}
