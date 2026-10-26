package main

import "fmt"

func main() {
	fmt.Println("=== 08. Monotonic Stack / Daily Temperatures ===")
	fmt.Println()

	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("температуры:", temps)
	fmt.Println("ждать дней: ", DailyTemperatures(temps)) // [1 1 4 2 1 1 0 0]
	fmt.Println("сумма ожидания:", WarmDays(temps))
	fmt.Println()

	fmt.Println("всё холодает:", DailyTemperatures([]int{9, 8, 7})) // [0 0 0]
	fmt.Println("один день:   ", DailyTemperatures([]int{5}))       // [0]
}
