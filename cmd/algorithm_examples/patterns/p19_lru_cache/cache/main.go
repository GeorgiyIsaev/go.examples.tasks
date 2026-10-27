package main

import "fmt"

func main() {
	fmt.Println("=== 19. LRU Cache (Least Recently Used) ===")
	fmt.Println()

	lru := NewLRUCache(3)

	fmt.Println("Последовательные Put/Get.")
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	probe := func(k int) {
		v := lru.Get(k)
		fmt.Printf("  Get(%d) = %d\n", k, v)
	}
	probe(1) // вернёт 1, узел 1 — самый recent
	probe(2) // 2 — recent
	probe(3) // 3 — recent

	fmt.Println()
	fmt.Println("Вытеснение: приходится выбрасывать.")
	lru.Put(4, 4) // LRU — 1 (после seq: Put(1),Put(2),Put(3),
	// probe(1),probe(2),probe(3) порядок recent=[3,2,1];
	// tail.prev = 1 → вытолкнем 1).
	probe(1) // -1
	probe(2)
	probe(3)
	probe(4)

	fmt.Println()
	fmt.Println("Обновление существующего ключа.")
	lru.Put(2, 20)
	probe(2) // 20

	fmt.Println()
	fmt.Println("Кэш с ёмкостью 0.")
	empty := NewLRUCache(0)
	empty.Put(10, 10)
	fmt.Println("  Get(10) =", empty.Get(10), "(ожидается -1)")

	fmt.Println()
	fmt.Println("Наглядное состояние кэша (не API, только для лекции):")
	printLRU("lru после всех операций", lru)
}

// Для наглядности: печать содержимого в порядке от recent к oldest.
func printLRU(name string, c *LRUCache) {
	fmt.Printf("  // %s: ", name)
	var first bool
	for n := c.head.next; n != c.tail; n = n.next {
		if first {
			fmt.Print(" <- ")
		}
		first = true
		fmt.Printf("%d:%d", n.key, n.value)
	}
	fmt.Println()
}
