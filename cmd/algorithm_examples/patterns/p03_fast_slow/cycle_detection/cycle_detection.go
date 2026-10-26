package main

// Паттерн: обнаружение цикла (Floyd's Cycle Detection, «черепаха и заяц»).
//
// Идея: два указателя. slow двигается на 1 шаг, fast — на 2.
// Если цикла нет — fast дойдёт до конца первым. Если цикл есть —
// fast и slow обязательно встретятся внутри цикла.
//
// Почему работает: в цикле fast «догоняет» slow со скоростью 1 шаг
// за итерацию. Разница длин сокращается на 1 каждый шаг, значит
// встреча гарантирована не более чем через длину цикла итераций.
//
// Когда применять:
//   - «есть ли цикл в связном списке»;
//   - «найти начало цикла»;
//   - «найти дубликат в массиве [1..n]» (Floyd's на индексах);
//   - «happy number» (Floyd's на числах);
//   - любая задача про повторяющееся состояние.
//
// Сложность: O(n) время, O(1) память.

// HasCycle — есть ли цикл в списке.
//
// Пример: 3 -> 2 -> 0 -> -4 -> 2 (обратно), вернёт true.
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// DetectCycle — узел, с которого начинается цикл. nil, если цикла нет.
//
// Идея: после встречи внутри цикла сбрасываем slow на head и
// двигаем оба указателя со скоростью 1. Точка встречи — начало цикла.
//
// Доказательство: пусть a — путь до цикла, b — от начала цикла до
// точки встречи, c — длина цикла. К моменту встречи slow прошёл a+b,
// fast прошёл a+b+kc. Разница = kc = путь slow (fast в 2 раза быстрее).
// Отсюда a = kc - b, то есть путь от head до цикла равен пути от
// точки встречи до начала цикла по циклу.
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// Фаза 1: ищем точку встречи внутри цикла.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Фаза 2: ищем начало цикла.
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// FindDuplicate — найти единственный дубликат в массиве.
// Условие: длина n+1, все числа в [1, n], ровно одно число повторяется
// (может — дважды и более раз).
// Массив менять нельзя, дополнительную память использовать нельзя.
//
// Идея: рассматриваем массив как связный список, где next(i) = nums[i].
// Тогда задача сводится к поиску начала цикла в этом «списке».
//
// Пример: [1,3,4,2,2] -> 2.
//
//	[3,1,3,4,2] -> 3.
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// Фаза 1: ищем встречу.
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Фаза 2: ищем вход в цикл.
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
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
