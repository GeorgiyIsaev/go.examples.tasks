package main

// Найти пару с суммой = target в отсортированном массиве
func twoSumSorted(nums []int, target int) (int, int, bool) {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		switch {
		case sum == target:
			return l, r, true
		case sum < target:
			l++ // нужно больше → двигаем левый вправо
		default:
			r-- // нужно меньше → двигаем правый влево
		}
	}
	return -1, -1, false
}
