package main

import "fmt"

func main() {
	fmt.Println("=== 07. Stack / History (история состояний) ===")
	fmt.Println()

	st := NewMinStack()
	for _, v := range []int{5, 3, 7, 1, 4} {
		st.Push(v)
		fmt.Printf("push %d -> min=%d top=%d len=%d\n", v, st.GetMin(), st.Top(), st.Len())
	}
	fmt.Println()
	for st.Len() > 0 {
		st.Pop()
		if st.Len() > 0 {
			fmt.Printf("pop   -> min=%d top=%d\n", st.GetMin(), st.Top())
		}
	}
	fmt.Println()

	fmt.Println("BackspaceCompare(\"ab#c\", \"ad#c\") =", BackspaceCompare("ab#c", "ad#c")) // true
	fmt.Println("BackspaceCompare(\"ab##\", \"a#b#\") =", BackspaceCompare("ab##", "a#b#")) // true
	fmt.Println("BackspaceCompare(\"a#c\", \"b\")     =", BackspaceCompare("a#c", "b"))     // false
}
