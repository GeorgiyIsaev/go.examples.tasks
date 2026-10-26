package main

import "fmt"

func main() {
	fmt.Println("=== 07. Stack / RPN (обратная польская нотация) ===")
	fmt.Println()

	fmt.Println("EvalRPN([2 1 + 3 *]) =", EvalRPN([]string{"2", "1", "+", "3", "*"}))   // 9
	fmt.Println("EvalRPN([4 13 5 / +]) =", EvalRPN([]string{"4", "13", "5", "/", "+"})) // 6
	fmt.Println("EvalRPN([10 6 9 3 + -11 * / * 17 + 5 +]) =",
		EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // 22
	fmt.Println()

	fmt.Println("CalcStringII(\"3+2*2\")   =", CalcStringII("3+2*2"))       // 7
	fmt.Println("CalcStringII(\" 3/2 \")   =", CalcStringII(" 3/2 "))       // 1
	fmt.Println("CalcStringII(\" 3+5 / 2 \") =", CalcStringII(" 3+5 / 2 ")) // 5
}
