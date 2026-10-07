package main

import "fmt"

//1. Удалить дубликаты с сохранением порядка
//Условие: Дан []int. Вернуть новый срез без повторов, сохранив порядок первых вхождений.
//Пример: [1,2,1,3,2,4] -> [1,2,3,4]
//Сигнатура:
//func Dedup(nums []int) []int

func Dedup(nums []int) []int {
	keys := make(map[int]struct{})
	result := make([]int, 0, len(nums))
	for _, num := range nums {
		if _, ok := keys[num]; !ok {
			keys[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 4}
	fmt.Println(nums)

	result := Dedup(nums)
	fmt.Println(result)

	result2 := Dedup(nil)
	fmt.Println(result2)

	nums3 := []int{1, 1, 1, 1, 1, 1}
	result3 := Dedup(nums3)
	fmt.Println(result3)

}
