package main

// Паттерн: разворот ПАРАМИ и группами — тот же трёхуказательный
// разворот, применённый точечно.
//
// SwapPairs (LeetCode 24): разворачиваем каждую пару соседних узлов.
// Главное — висеть на узле ДО пары (dummy/prev), иначе после разворота
// потеряется связь с началом списка.
//
// ReverseKGroup (LeetCode 25): то же, но блоками по k. Если хвост
// короче k — его не трогаем. Проверка «осталось k узлов» — вместо
// того чтобы считать длину заранее.
//
// Сложность: O(n) время, O(1) память (итеративно).

// SwapPairs — перевернуть каждую пару соседних узлов.
// Пример: 1->2->3->4->5 => 2->1->4->3->5.
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next      // первый в паре
		b := prev.Next.Next // второй в паре

		// Перекидываем стрелки: prev -> b -> a -> (бывший b.Next).
		prev.Next = b
		a.Next = b.Next
		b.Next = a

		prev = a // теперь prev — последний узел обработанной пары
	}
	return dummy.Next
}

// ReverseKGroup — перевернуть списки по k узлов (k >= 1).
// Хвост короче k остаётся как есть.
// Пример: 1->2->3->4->5, k=2 => 2->1->4->3->5; k=3 => 3->2->1->4->5.
func ReverseKGroup(head *ListNode, k int) *ListNode {
	// Проверим, что впереди ровно k узлов.
	hasK := func(from *ListNode) bool {
		for i := 0; i < k; i++ {
			if from == nil {
				return false
			}
			from = from.Next
		}
		return true
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for hasK(prev.Next) {
		// Разворачиваем группу, вставляя каждый её узел сразу
		// после prev: порядок переворачивается автоматически.
		// Первый узел уже стоит на месте (i=0 — ничего не делаем).
		groupStart := prev.Next
		cur := groupStart
		for i := 0; i < k; i++ {
			next := cur.Next // запомнили, куда идём дальше
			if i > 0 {
				cur.Next = prev.Next // ставим в начало развёрнутого
				prev.Next = cur
			}
			cur = next
		}
		// groupStart теперь последний в группе — подшиваем остаток.
		groupStart.Next = cur
		prev = groupStart // он же хвост обработанной группы
	}
	return dummy.Next
}
