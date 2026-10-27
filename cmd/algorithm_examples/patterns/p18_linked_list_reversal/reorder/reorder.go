package main

// Reorder List (LeetCode 143) — «склейка» списка в виде
// L0→Ln→L1→Ln-1→L2→Ln-2→...
//
// Три шага, каждый уже знаком из этого паттерна:
//  1. найти середину — fast/slow (p02);
//  2. развернуть вторую половину — ReverseList (этот же паттерн);
//  3. склеить по одному узлу из каждой половины — merge как в
//     «сжатии» списка.
//
// Сложность: O(n) время, O(1) память (без дополнительных структур).

// ReorderList перестраивает список на месте.
//
// Пример: 1->2->3->4 => 1->4->2->3;
// 1->2->3->4->5 => 1->5->2->4->3.
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. Середина: slow останавливается на первом (для чётных)
	//    или центральном (для нечётных) узле.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Разворачиваем вторую половину, начиная со slow.Next.
	second := slow.Next
	slow.Next = nil // разрываем список на две половины
	var prev *ListNode
	cur := second
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// prev — голова развёрнутой второй половины.

	// 3. Склеиваем: first -> prev -> first.Next -> ...
	first := head
	secondHead := prev
	for secondHead != nil {
		n1, n2 := first.Next, secondHead.Next
		first.Next = secondHead
		secondHead.Next = n1
		first = n1
		secondHead = n2
	}
}
