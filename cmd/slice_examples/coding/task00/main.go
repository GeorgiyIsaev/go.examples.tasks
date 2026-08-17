package main

import "fmt"

// Написать функция merge
func merge(a, b []int) []int {
	size := len(a) + len(b)
	result := make([]int, 0, size)

	tempA, tempB := 0, 0

	for tempA < len(a) && tempB < len(b) {

		if b[tempB] <= a[tempA] {
			result = append(result, b[tempB])
			tempB++
		} else {
			result = append(result, a[tempA])
			tempA++
		}

	}

	result = append(result, a[tempA:]...)
	result = append(result, b[tempB:]...)
	return result

}

func main() {
	fmt.Println(merge([]int{1, 2, 5}, []int{3, 4})) // [1 2 3 4 5]
	fmt.Println(merge([]int{}, []int{1, 2}))        // [1 2]
	fmt.Println(merge([]int{1, 2}, []int{}))        // [1 2]
	fmt.Println(merge([]int{1, 2, 2}, []int{2, 3})) // [1 2 2 2 3]
}
