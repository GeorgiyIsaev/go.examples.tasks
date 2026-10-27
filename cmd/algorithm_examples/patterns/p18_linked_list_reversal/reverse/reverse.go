package main

// Паттерн: разворот односвязного списка.
//
// Ключевая идея: пока не «перекинули» указатель, список рвётся —
// поэтому работаем через ТРИ указателя (prev, cur, next):
//
//	prev, cur := (*ListNode)(nil), head
//	for cur != nil {
//	    next := cur.Next   // 1. запомнили, куда дальше
//	    cur.Next = prev    // 2. развернули указатель
//	    prev = cur         // 3. шагнули вперёд
//	    cur = next
//	}
//	return prev           // новая голова = бывший хвост
//
// Сложность: O(n) время, O(1) память (итеративно).
// Рекурсивный вариант — O(n) памяти на стек вызовов: короче код,
// но глубина = длина списка.

// ReverseList — развернуть список целиком (LeetCode 206).
// Пример: 1->2->3->4->5  =>  5->4->3->2->1.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // запомнили хвост
		cur.Next = prev  // развернули стрелку
		prev = cur       // шагнули
		cur = next
	}
	return prev
}

// ReverseListRecursive — тот же разворот рекурсией: сначала развернуть
// хвост, потом повернуть текущий узел. Код короче, памяти O(n).
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseListRecursive(head.Next)
	head.Next.Next = head // бывший следующий теперь указывает на нас
	head.Next = nil       // а мы больше не указываем вперёд
	return newHead
}

// ReverseBetween — развернуть список между позициями left и right
// включительно (LeetCode 92, позиции 1-based).
//
// Трюк: ставим «заглушку» перед head и доходим до узла left-1 —
// дальше classic-разворот ровно (right-left+1) узлов, хвост
// подшиваем к началу не тронутой части.
//
// Пример: 1->2->3->4->5, left=2, right=4 => 1->4->3->2->5.
func ReverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	// Разворачиваем (right-left+1) узлов, начиная с cur.
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}
