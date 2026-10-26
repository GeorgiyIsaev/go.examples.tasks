package main

import "fmt"

func main() {
	fmt.Println("=== 07. Stack / Brackets (скобки и распаковка) ===")
	fmt.Println()

	for _, s := range []string{"([]{})", "([)]", "(", "}"} {
		fmt.Printf("IsValid(%q) = %v\n", s, IsValid(s))
	}
	fmt.Println()

	for _, s := range []string{"3[a2[c]]", "2[abc]3[cd]ef", "abc3[cd]xyz"} {
		fmt.Printf("DecodeString(%q) = %q\n", s, DecodeString(s))
	}
}
