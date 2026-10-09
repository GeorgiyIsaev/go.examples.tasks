package main

import "fmt"

//12. Подмассив с суммой K
//Условие:
//Дан []int (могут быть отрицательные) и число k.
//Вернуть количество непрерывных подмассивов,
//сумма элементов которых равна k.
//Примеры:
//•	nums = [1,1,1], k = 2 → 2 (подмассивы [0:1] и [1:2])
//•	nums = [1,2,3], k = 3 → 2 ([1,2] и [3])
//•	nums = [1,-1,0], k = 0 → 3 ([1,-1], [0], [1,-1,0])
//•	nums = [], k = 0 → 0
//Сигнатура:
//func SubarraySum(nums []int, k int) int

func SubarraySum(nums []int, k int) int {
	//Алгоритм
	//- Заводим map[int]int — сколько раз каждая префиксная
	//		сумма встречалась. Идём по массиву, накапливая currentSum:
	//- Для каждого элемента прибавляем его к currentSum.
	//- Проверяем: сколько раз в map встречалось currentSum - k?
	//		Прибавляем это число к count.
	//- Записываем currentSum в map (увеличиваем счётчик на 1).

	prefixCount := map[int]int{0: 1}
	count := 0
	currentSum := 0

	for _, num := range nums {
		currentSum += num

		//Сколько раз ранне встречалась сумма
		if prevCount, ok := prefixCount[currentSum-k]; ok {
			count += prevCount
		}
		prefixCount[currentSum]++
	}
	return count
}

func main() {
	fmt.Println(SubarraySum([]int{1, 1, 1}, 2))  // 2
	fmt.Println(SubarraySum([]int{1, 2, 3}, 3))  // 2
	fmt.Println(SubarraySum([]int{1, -1, 0}, 0)) // 3
	fmt.Println(SubarraySum([]int{3}, 3))        // 1
	fmt.Println(SubarraySum([]int{}, 0))         // 0
	fmt.Println(SubarraySum([]int{1}, 0))        // 0
}
