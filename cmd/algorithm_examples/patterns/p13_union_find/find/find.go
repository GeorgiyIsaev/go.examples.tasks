package main

// Паттерн: Union Find (DSU) — «связность за почти O(1): find/union,
// ранг, сжатие пути, счётчик компонент».
//
// Ключевая идея: каждый узел знает своего родителя; корень знает
// размер своей компоненты (parent[x] < 0 → x — корень, -parent[x] — размер).
// При union меньший корень подчиняется большему (по модулю размера),
// при find — сжатие пути (ребёнок «припоминает» найденный корень).
//
// Сложность: find + union — O(α(n)) амортизированно, память — O(n).
// α(n) ≤ 4 для любого реального n; на собесе это «почти константа».

// DSU держит компоненты; parent[x] < 0 ⇒ x — корень со размером -parent[x].
type DSU struct{ parent []int }

// NewDSU создаёт n изолированных компонент.
func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = -1
	}
	return &DSU{parent: p}
}

// Find возвращает корень компоненты x с сжатием пути.
func (d *DSU) Find(x int) int {
	if d.parent[x] < 0 {
		return x
	}
	d.parent[x] = d.Find(d.parent[x])
	return d.parent[x]
}

// Union соединяет компоненты a и b; возвращает true, если соединение новое.
func (d *DSU) Union(a, b int) bool {
	a, b = d.Find(a), d.Find(b)
	if a == b {
		return false
	}
	// меньший размер подчиняется большему.
	if -d.parent[a] < -d.parent[b] {
		a, b = b, a
	}
	d.parent[a] += d.parent[b] // обновить размер большего
	d.parent[b] = a
	return true
}

// Same — находятся ли a и b в одной компоненте.
func (d *DSU) Same(a, b int) bool { return d.Find(a) == d.Find(b) }

// Size — размер компоненты, содержащей x (только для валидного x).
func (d *DSU) Size(x int) int { return -d.parent[d.Find(x)] }

// Components — текущее количество компонент.
func (d *DSU) Components() int {
	n := 0
	for _, p := range d.parent {
		if p < 0 {
			n++
		}
	}
	return n
}
