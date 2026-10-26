package main

import "fmt"

func main() {
	fmt.Println("=== 05. Hash Map / Anagrams (анаграммы и группы) ===")
	fmt.Println()

	fmt.Println("IsAnagram(\"anagram\", \"nagaram\") =", IsAnagram("anagram", "nagaram")) // true
	fmt.Println("IsAnagram(\"rat\", \"car\")       =", IsAnagram("rat", "car"))           // false
	fmt.Println("IsAnagram(\"ab\", \"a\")          =", IsAnagram("ab", "a"))              // false
	fmt.Println()

	groups := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println("GroupAnagrams([eat,tea,tan,ate,nat,bat]):")
	for _, g := range groups {
		fmt.Println(" ", g)
	}
	fmt.Println()

	fmt.Println("GroupAnagrams([]) =", GroupAnagrams([]string{}))
}
