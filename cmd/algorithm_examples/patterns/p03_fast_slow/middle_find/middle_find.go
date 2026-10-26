package main

// Паттерн: fast/slow для поиска середины и связанных задач.
//
// Ключевая идея та же: fast двигается в 2 раза быстрее. Когда fast
// доходит до конца, slow оказывается ровно в середине.
//
// Варианты использования:
//   - найти середину списка (для палиндрома, merge sort);
//   - проверить список на палиндром;
//   - разделить список пополам;
//   - любая задача, где нужно «дойти до половины, не зная длины».
//
// Сложность: O(n) время, O(1) память.

// MiddleNode — средний узел списка.
// Если узлов чётное количество, возвращает ВТОРОЙ средний (как на LeetCode).
//
// Пример: [1,2,3,4,5] -> узел 3.
//
//	[1,2,3,4]   -> узел 3 (второй средний).
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// IsPalindromeList — проверка, является ли связный список палиндромом.
//
// Идея: находим середину (slow), разворачиваем вторую половину,
// сравниваем с первой, при желании — восстанавливаем.
//
// Пример: [1,2,2,1] -> true. [1,2] -> false.
func IsPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 1. Находим середину.
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Разворачиваем вторую половину.
	prev := (*ListNode)(nil)
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// prev теперь голова развёрнутой второй половины.

	// 3. Сравниваем первую половину с развёрнутой второй.
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
