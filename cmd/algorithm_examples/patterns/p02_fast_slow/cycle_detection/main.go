package main

import "fmt"

// Запуск: go run ./cycle_detection
func main() {
	fmt.Println("=== 02. Fast / Slow / Cycle Detection (Floyd's) ===")
	fmt.Println()

	// Список без цикла: 1 -> 2 -> 3 -> 4
	noCycle := BuildList([]int{1, 2, 3, 4})
	fmt.Printf("HasCycle(1->2->3->4)                 = %v\n", HasCycle(noCycle))
	fmt.Printf("DetectCycle(1->2->3->4)              = %v\n", DetectCycle(noCycle))
	fmt.Printf("CountCycleLength(1->2->3->4)         = %d\n", CountCycleLength(noCycle))

	// Список с циклом: 3 -> 2 -> 0 -> -4 -> (обратно на 2)
	cyc := BuildList([]int{3, 2, 0, -4})
	tail := findNodeByIndex(cyc, 3)
	entry := findNodeByIndex(cyc, 1) // узел со значением 2
	tail.Next = entry

	fmt.Printf("HasCycle(3->2->0->-4->2...)          = %v\n", HasCycle(cyc))
	fmt.Printf("DetectCycle(...).Val                 = %d (ожидается 2)\n", DetectCycle(cyc).Val)
	fmt.Printf("CountCycleLength(...)                = %d (ожидается 3)\n", CountCycleLength(cyc))

	fmt.Println()

	fmt.Printf("FindDuplicate([1,3,4,2,2])           = %d\n", FindDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("FindDuplicate([3,1,3,4,2])           = %d\n", FindDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("FindDuplicate([1,1])                 = %d\n", FindDuplicate([]int{1, 1}))
}
