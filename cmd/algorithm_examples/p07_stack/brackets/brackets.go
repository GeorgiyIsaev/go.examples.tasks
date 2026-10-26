package main

import (
	"strings"
)

// Паттерн: стек — «последний зашёл, первый вышел» (LIFO).
//
// Когда стек: встречаем открывающую конструкцию — кладём контекст
// (незакрытую скобку, «что ждёт продолжения»); встречаем закрывающую
// — сверяем/разворачиваем с вершиной. Стек хранит именно то, что было
// «до» текущей позиции и ещё не закрыто.
//
// Сложность: O(n) времени, O(n) памяти в худшем случае.

// IsValid — правильная ли скобочная последовательность (LeetCode 20).
// Стек закрывающих скобок: каждая закрывающая должна совпасть
// с последней не закрытой открывающей.
//
// Пример: "([]{})" -> true, "([)]" -> false, "(" -> false.
func IsValid(s string) bool {
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}

	for _, r := range s {
		switch {
		case r == '(' || r == '[' || r == '{':
			stack = append(stack, r) // открыли — ждём закрытия
		default:
			closing, isClosing := pairs[r]
			if !isClosing {
				return false // посторонний символ
			}
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false // не с кем закрылось / чужая скобка
			}
			stack = stack[:len(stack)-1] // закрыли — сняли
		}
	}
	return len(stack) == 0 // всё закрыто?
}

// DecodeString — распаковать строку вида "3[a2[c]]" (LeetCode 394).
// "3[a2[c]]" -> "accaccacc".
//
// Идея: два стека (или один стек пар) — «контексты»:
//   - встретили цифру → запомнили число на будущее;
//   - встретили '[' → положили в стек то, что НАКОПИЛИ ДО скобки,
//     и счётчик повторов; начинаем копить заново;
//   - встретили ']' → вершина стека + (текущее накопленное × count)
//     становится новым накопленным.
//
// Вложенные скобки работают, потому что каждый '[' кладёт новый слой.
func DecodeString(s string) string {
	countStack := []int{}
	strStack := []string{}
	current := strings.Builder{}

	num := 0 // многозначные числа: 12[a]
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
			num = num*10 + int(r-'0')
		case r == '[':
			// Открываем новый слой: сохранить накопленное и число.
			countStack = append(countStack, num)
			strStack = append(strStack, current.String())
			current.Reset()
			num = 0
		case r == ']':
			// Закрываем: предыдущее + текущее × count.
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			repeated := strings.Repeat(current.String(), count)
			current.Reset()
			current.WriteString(prev)
			current.WriteString(repeated)
		default:
			current.WriteRune(r)
		}
	}
	return current.String()
}
