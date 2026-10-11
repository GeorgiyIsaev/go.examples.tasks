package main

// Паттерн: fast/slow для поиска середины и связанных задач.
//
// Ключевая идея та же: fast двигается в 2 раза быстрее. Когда fast
// доходит до конца, slow оказывается ровно в середине.
//
// Варианты использования:
//   - найти середину списка (для палиндрома, merge sort);
//   - найти «happy number» (Floyd's на числах);
//   - определить «счастливое число»;
//   - любая задача, где нужно «дойти до половины, не зная длины».

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

// IsHappy — является ли число «счастливым».
// Счастливое число: сумма квадратов цифр, потом снова, и так пока
// не получим 1 (счастливое) или не зациклимся (несчастливое).
//
// Пример: 19 -> true (1+81=82, 64+4=68, 36+64=100, 1+0+0=1).
//
//	2  -> false (зациклится на 4-16-37-58-89-145-42-20-4).
func IsHappy(n int) bool {
	next := func(x int) int {
		sum := 0
		for x > 0 {
			d := x % 10
			sum += d * d
			x /= 10
		}
		return sum
	}

	slow, fast := n, n
	for {
		slow = next(slow)
		fast = next(next(fast))
		if slow == fast {
			break
		}
	}
	return slow == 1
}

// CountCycleLength — длина цикла в списке (если он есть).
// 0, если цикла нет. Полезно для задач, где нужно знать размер цикла.
func CountCycleLength(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// Нашли точку внутри цикла. Считаем длину.
			length := 1
			for cur := slow.Next; cur != slow; cur = cur.Next {
				length++
			}
			return length
		}
	}
	return 0
}
