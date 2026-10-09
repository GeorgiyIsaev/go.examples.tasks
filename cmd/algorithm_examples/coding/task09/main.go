package main

import "fmt"

//9. Произведение всех элементов кроме самого себя
//Условие: Дан []int. Вернуть новый срез, где result[i] = произведение
//всех элементов исходного среза, кроме nums[i].
//Деление использовать нельзя. Работать за O(n).
//Примеры:
//•	[1,2,3,4] → [24,12,8,6]
//•	[-1,1,0,-3,3] → [0,0,9,0,0]
//•	[5] → [1]
//Сигнатура:
//func ProductExceptSelf(nums []int) []int

func ProductExceptSelf_false(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				result[i] = nums[j] * result[i]
			}

		}
	}
	//Квадратичное решение  O(n²)
	return result
}

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Проход 1: result[i] = произведение всего слева от i
	//Первфикс
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Проход 2: домножаем на произведение всего справа от i
	//Суфикс
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func main() {
	cl := []int{1, 2, 3, 4}
	fmt.Println(ProductExceptSelf(cl))

	cl = []int{5}
	fmt.Println(ProductExceptSelf(cl))
}
