package main

import "fmt"

func main() {
	fmt.Println("=== 13. Union Find (связность и компоненты) ===")
	fmt.Println()

	// Пример 1: классическое «соединяй ребра, считай компоненты».
	fmt.Println("Пример 1: связность рёбер.")
	n := 6
	d := NewDSU(n)
	edges := [][2]int{
		{0, 1}, {1, 2}, // компонента A: {0,1,2}
		{3, 4}, // компонента B: {3,4}
		// 5 — один
	}
	for _, e := range edges {
		if d.Union(e[0], e[1]) {
			fmt.Printf("  объединил %d-%d\n", e[0], e[1])
		} else {
			fmt.Printf("  %d-%d уже в одной компоненте\n", e[0], e[1])
		}
	}
	fmt.Println("  компонент осталось:", d.Components(), "(ожидается 3: {0,1,2}, {3,4}, {5})")
	fmt.Println("  размер компоненты 0:", d.Size(0))
	fmt.Println("  размер компоненты 3:", d.Size(3))

	// Пример 2: «в одной ли группе» через Same.
	fmt.Println()
	fmt.Println("Пример 2: Same/Find.")
	fmt.Println("  0 и 2 вместе? ->", d.Same(0, 2), "(true)")
	fmt.Println("  0 и 3 вместе? ->", d.Same(0, 3), "(false)")
	fmt.Println("  корень 4:", d.Find(4), "(ожидается 3)")

	// Пример 3: ранг — меньший подчиняется большему; дерево остаётся широким.
	fmt.Println()
	fmt.Println("Пример 3: ранг (меньший ранг подчиняется большему).")
	d2 := NewDSU(4)
	d2.Union(0, 1)
	d2.Union(2, 3)
	d2.Union(1, 2)
	fmt.Println("  после объединений 0-1, 2-3, 1-2:")
	fmt.Println("    родитель 0:", d2.Find(0), "размер:", -d2.parent[d2.Find(0)])
	fmt.Println("    компонент:", d2.Components(), "(ожидается 1)")

	// Пример 4: рёбра, которые ничего не меняют (петли/дубликаты).
	fmt.Println()
	fmt.Println("Пример 4: Union уже соединённых.")
	fmt.Println("  Union(0,0) ->", d2.Union(0, 0), "(false)")
	fmt.Println("  Union(2,0) ->", d2.Union(2, 0), "(false, уже вместе)")
}
