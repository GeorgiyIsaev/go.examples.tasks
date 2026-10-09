package main

import (
	"fmt"
	"sort"
)

//Сумма трёх = 0 (Three Sum)
//Условие: Дан []int.
//Вернуть все уникальные тройки [a, b, c] таких,
//что a + b + c = 0. Порядок троек не важен,
//значения внутри тройки — по возрастанию,
//повторов троек быть не должно.
//Примеры:
//•	[-1,0,1,2,-1,-4] → [[-1,-1,2],[-1,0,1]]
//•	[0,0,0] → [[0,0,0]]
//•	[1,2,3] → []
//•	[] → []
//Сигнатура:
//func ThreeSum(nums []int) [][]int

func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	//1 Сортировка
	sort.Ints(nums)
	n := len(nums)

	//2. Фиксируем первый элемент
	for i := 0; i < n-2; i++ {
		//Проверка на повтор числа
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//Оптимизация: Если уже больше 0 нет смысла проверять
		if nums[i] > 0 {
			break
		}

		//3. Два указатля
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// нашли тройку
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Дедупликация уровня left/right:
				// сдвигаем left, пока следующее значение такое же
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// сдвигаем right, пока предыдущее значение такое же
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// и сдвигаем оба указателя к новым значениям
				left++
				right--
			}
		}
	}

	//1. Сортировка
	//2. Фиксируем i
	//   2a. Skip дубликаты на уровне i
	//   2b. Two Sum на оставшемся подмассиве (два указателя)
	//       - Skip дубликаты на уровне left/right
	//3. Возврат
	return result
}

func main() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(ThreeSum([]int{0, 0, 0}))             // [[0 0 0]]
	fmt.Println(ThreeSum([]int{1, 2, 3}))             // []
	fmt.Println(ThreeSum([]int{}))                    // []
}
