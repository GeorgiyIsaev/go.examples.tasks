package main

// ListNode — узел односвязного списка (классическое определение LeetCode).
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList строит список из среза, возвращает голову.
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

// ListToSlice разворачивает список в срез (для тестов и демо).
// Для списков с циклом зациклится — здесь их нет.
func ListToSlice(head *ListNode) []int {
	result := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return result
}
