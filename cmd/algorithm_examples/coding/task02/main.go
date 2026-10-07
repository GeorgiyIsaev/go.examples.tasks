package main

import "fmt"

//2. Two Sum
//Условие: Дан []int и target. Вернуть индексы двух чисел, которые в сумме дают target. Ровно одно решение.
//Пример: nums = [2,7,11,15], target = 9 -> [0,1]
//Сигнатура:
//func TwoSum(nums []int, target int) []int

func TwoSum_bad(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	//Плохое решение - квадратичная сложность

	return result
}

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) //Карта Значение -> индекс
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{i, j}
		}
		seen[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}

	fmt.Println(nums)

	target := 9
	fmt.Println("Сумма", target, "", TwoSum(nums, target))

	target = 26
	fmt.Println("Сумма", target, "", TwoSum(nums, target))

	target = 17
	fmt.Println("Сумма", target, "", TwoSum(nums, target))

	target = 1
	fmt.Println("Сумма", target, "", TwoSum(nums, target))

}
