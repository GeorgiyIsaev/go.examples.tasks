package main

import "fmt"

func main() {
	fmt.Println("=== 06. Trie / Basic (вставка, поиск, автодополнение) ===")
	fmt.Println()

	tr := NewTrie()
	for _, w := range []string{"apple", "app", "apricot", "banana", "band"} {
		tr.Insert(w)
	}

	fmt.Println("Search(\"apple\")   =", tr.Search("apple"))   // true
	fmt.Println("Search(\"app\")     =", tr.Search("app"))     // true (вставили)
	fmt.Println("Search(\"appl\")    =", tr.Search("appl"))    // false — только префикс
	fmt.Println("Search(\"bananaa\") =", tr.Search("bananaa")) // false
	fmt.Println()

	fmt.Println("StartsWith(\"ap\")     =", tr.StartsWith("ap"))     // true
	fmt.Println("StartsWith(\"appL\")   =", tr.StartsWith("appl"))   // true
	fmt.Println("StartsWith(\"banana\") =", tr.StartsWith("banana")) // true
	fmt.Println("StartsWith(\"cat\")    =", tr.StartsWith("cat"))    // false
	fmt.Println()

	fmt.Println("WordsWithPrefix(\"ap\")  =", tr.WordsWithPrefix("ap"))  // [app apple apricot]
	fmt.Println("WordsWithPrefix(\"ban\") =", tr.WordsWithPrefix("ban")) // [banana band]
	fmt.Println("WordsWithPrefix(\"\")    =", tr.WordsWithPrefix(""))    // все слова по алфавиту
	fmt.Println("WordsWithPrefix(\"z\")   =", tr.WordsWithPrefix("z"))   // []
	fmt.Println()

	// Повторная вставка ничего не ломает.
	tr.Insert("apple")
	fmt.Println("после повторного Insert(\"apple\"): Search =", tr.Search("apple")) // true
}
