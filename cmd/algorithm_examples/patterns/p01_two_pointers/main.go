package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 6, 8, 11}
	target := 10

	i, j, ok := twoSumSorted(nums, target)
	if !ok {
		fmt.Printf("Пара с суммой %d не найдена\n", target)
		return
	}

	fmt.Printf("nums = %v, target = %d\n", nums, target)
	fmt.Printf("Индексы: %d и %d\n", i, j)
	fmt.Printf("Значения: %d + %d = %d\n", nums[i], nums[j], nums[i]+nums[j])
}
