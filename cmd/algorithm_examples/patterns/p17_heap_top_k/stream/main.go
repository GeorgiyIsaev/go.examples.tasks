package main

import "fmt"

func main() {
	fmt.Println("=== 17. Heap / Stream (медиана и скользящее среднее) ===")
	fmt.Println()

	m := NewMedianFinder()
	fmt.Println("поток: 6 10 2 6 5 4 0 6 3 8")
	fmt.Print("медианы после каждого: ")
	for _, v := range []int{6, 10, 2, 6, 5, 4, 0, 6, 3, 8} {
		m.AddNum(v)
		fmt.Printf("%.1f ", m.FindMedian())
	}
	fmt.Println()
	fmt.Println("(после 6 -> 6.0; после 6,10 -> 8.0; после +2 -> 6.0)")
	fmt.Println()

	ma := NewMovingAverage(3)
	fmt.Println("MovingAverage(size=3), значения 1 3 5 7 9:")
	for _, v := range []int{1, 3, 5, 7, 9} {
		fmt.Printf("  Next(%d) = %.4f\n", v, ma.Next(v))
	}
}
