package main

// ListNode — узел односвязного списка.
// Классическое определение из LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList строит список из среза. Нужно для тестов и демо.
// Возвращает голову списка. Для пустого среза — nil.
func BuildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

// ListToSlice разворачивает список в срез.
// ВНИМАНИЕ: для списка с циклом зациклится! Использовать только там,
// где цикла гарантированно нет.
func ListToSlice(head *ListNode) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}

// findNodeByIndex возвращает узел по индексу (0-based).
// Нужно в тестах, чтобы создать цикл в конкретной позиции.
func findNodeByIndex(head *ListNode, idx int) *ListNode {
	cur := head
	for i := 0; i < idx && cur != nil; i++ {
		cur = cur.Next
	}
	return cur
}
