package main

// Daily Temperatures (LeetCode 739): для каждого дня — сколько дней
// ждать более тёплого. Классический «следующий больший» с РАССТОЯНИЕМ,
// а не со значением.
//
// Отличие от next_greater тем, что в ответе — ИНДЕКС ближайшего
// большего, а расстояние = i - stack_top. Монотонность та же:
// убывающие значения в стеке.
//
// Сложность: O(n) — каждый индекс входит и выходит один раз.
//
// Пример: [73,74,75,71,69,72,76,73] -> [1,1,4,2,1,1,0,0].

// DailyTemperatures возвращает, сколько дней ждать более тёплой погоды.
// Если warmer впереди нет — 0.
func DailyTemperatures(temps []int) []int {
	n := len(temps)
	answer := make([]int, n)

	stack := []int{} // индексы дней с убывающими температурами
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temps[i] > temps[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = i - top // ровно столько дней ждали
		}
		stack = append(stack, i)
	}
	// Оставшиеся в стеке — warmer так и не наступил: answer[i] = 0.
	return answer
}

// WarmDays — маленькая демонстрация: суммарное ожидание (для печати
// в main и как проверка, что «ноль в конце» работает).
func WarmDays(temps []int) int {
	total := 0
	for _, d := range DailyTemperatures(temps) {
		total += d
	}
	return total
}
