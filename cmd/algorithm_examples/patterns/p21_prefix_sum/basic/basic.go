package main

// Паттерн: префиксные суммы (prefix sums).
//
// Идея: заранее (за один проход) посчитать накопительные суммы —
// тогда сумма ЛЮБОГО подмассива [i..j] считается за O(1):
//
//	sum(i..j) = prefix[j+1] - prefix[i]   (prefix[0] = 0)
//
// Стоимость: O(n) preprocessing + O(1) на запрос. Когда в условии
// «много запросов на сумму диапазона» — это почти всегда оно.
//
// Продвинутая форма (p05 pairs, LC560): если вместо готового
// массива спрашивают «существует ли подмассив с суммой k», то
// prefix-суммы превращаются в карту встречаемости за O(n).

// RunningSum — бегущая сумма: result[i] = nums[0]+...+nums[i]
// (LeetCode 1480).
//
// Пример: [1,2,3,4] -> [1,3,6,10].
func RunningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		result[i] = sum
	}
	return result
}

// PivotIndex — индекс, где суммы слева и справа равны (LeetCode 724).
// Левая сумма по умолчанию 0; правая — total - running.
// Нет такого индекса -> -1.
//
// Пример: [1,7,3,6,5,6] -> 3 (11 == 11).
func PivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	left := 0
	for i, v := range nums {
		// right = total - left - v
		if left == total-left-v {
			return i
		}
		left += v
	}
	return -1
}

// RangeSumQuery — класс «суммы по диапазону после O(n) подготовки»
// (LeetCode 303: immutable, конструктор один раз, запросы O(1)).
type RangeSumQuery struct {
	prefix []int // prefix[i] = сумма nums[:i]
}

// NewRangeSumQuery строит префиксы за O(n).
func NewRangeSumQuery(nums []int) *RangeSumQuery {
	prefix := make([]int, len(nums)+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	return &RangeSumQuery{prefix: prefix}
}

// SumRange — сумма nums[left..right] включительно, O(1).
func (q *RangeSumQuery) SumRange(left, right int) int {
	return q.prefix[right+1] - q.prefix[left]
}
