// Паттерн: стек как «история состояний» — что было до этого.
//
// Ключевая идея: держать не только данные, но и производную от них
// информацию (минимум, максимум, «что было раньше») ВМЕСТЕ с каждым
// шагом. Тогда «спросить про прошлое» — чтение вершины, O(1).
//
// Типовые задачи:
//   - Min Stack (LC155): push/pop/top/getMin за O(1);
//   - Backspace String Compare (LC844): «стирание» — откат состояния;
//   - дальше в курсе: монотонный стек (p08) — история «предыдущий
//     больший/меньший».
package main

// MinStack — стек с O(1)-доступом к минимуму (LeetCode 155).
//
// Трюк: рядом с каждым значением храним минимум НА МОМЕНТ его
// записи. Min вершины — минимум «всего, что лежит ниже и включая
// себя», а общий минимум стека всегда равен минимуму вершины.
//
// Альтернатива — два стека (данные + монотонный стек минимумов),
// экономит память на дублях минимума; вариант с парами проще
// и читается лучше.
type MinStack struct {
	data []struct{ val, min int }
}

// NewMinStack — конструктор (в Go без него пришлось бы инициализировать
// поле снаружи).
func NewMinStack() *MinStack {
	return &MinStack{}
}

// Push кладёт элемент; минимум пересчитывается за O(1).
func (s *MinStack) Push(val int) {
	min := val
	if len(s.data) > 0 && s.data[len(s.data)-1].min < min {
		min = s.data[len(s.data)-1].min
	}
	s.data = append(s.data, struct{ val, min int }{val, min})
}

// Pop снимает верхний элемент.
func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

// Top — верхний элемент.
func (s *MinStack) Top() int {
	return s.data[len(s.data)-1].val
}

// GetMin — текущий минимум, O(1).
func (s *MinStack) GetMin() int {
	return s.data[len(s.data)-1].min
}

// Len — размер стека (удобно для тестов и проверок).
func (s *MinStack) Len() int {
	return len(s.data)
}

// BackspaceCompare — равны ли строки после удаления '#' как
// backspace (LeetCode 844).
//
// Два подхода:
//  1. стек-симуляция: проходим строку, '#' — pop, иначе push;
//     сравниваем стеки (или строим строки);
//  2. два указателя с конца: отсчитывать «сколько удалить»,
//     не выделяя память (см. ниже — реализован стек-вариант,
//     он прямолинее и переиспользует навык p07).
//
// Пример: "ab#c" / "ad#c" -> true (обе "ac").
func BackspaceCompare(s, t string) bool {
	return buildBackspaced(s) == buildBackspaced(t)
}

// buildBackspaced прогоняет строку через стек-симуляцию.
func buildBackspaced(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // стерли последний
			}
			continue
		}
		stack = append(stack, r)
	}
	return string(stack)
}
