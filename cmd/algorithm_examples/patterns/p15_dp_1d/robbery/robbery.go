package main

// House Robber — классика «грабим / пропускаем» (LeetCode 198 и 213).
//
// Состояние: dp[i] — максимум, который можно собрать среди первых
// i домов (границы i включительно рассматривать нельзя — соседи).
// Рекуррентность с двумя ветками — сердце паттерна:
//
//	dp[i] = max(dp[i-1],          // дом i ПРОПУСКАЕМ
//	            dp[i-2] + nums[i]) // дом i ГРАБИМ (сосед уже учтён в dp[i-2])
//
// База: dp[0] = nums[0], dp[1] = max(nums[0], nums[1]).
// Память: dp[i] зависит только от двух предыдущих -> две переменные.

// Rob — максимум суммы, не грабя два соседних дома (LeetCode 198).
//
// Пример: [1,2,3,1] -> 4 (1 + 3).
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev2 := 0 // dp[i-2]
	prev1 := 0 // dp[i-1], на старте «до первого дома» = 0
	for _, v := range nums {
		// новый максимум: не брать v или брать v + максимум до соседа
		cur := prev1
		if prev2+v > prev1 {
			cur = prev2 + v
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}

// RobCircle — дома стоят по КРУГУ: первый и последний соседи
// (LeetCode 213). Грабить оба сразу нельзя.
//
// Трюк «разрыва»: круг — это два линейных случая:
//   - не грабим первый дом  -> обычный Rob(nums[1:])
//   - не грабим последний   -> обычный Rob(nums[:n-1])
//
// Ответ — максимум из двух (или Rob целиком, если n <= 2...
// при n=2 просто max, при n=1 — nums[0]).
//
// Пример: [2,3,2] -> 3 (нельзя взять 2 и 2 по кругу).
func RobCircle(nums []int) int {
	n := len(nums)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return nums[0]
	case n == 2:
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	withoutFirst := Rob(nums[1:])
	withoutLast := Rob(nums[:n-1])
	if withoutFirst > withoutLast {
		return withoutFirst
	}
	return withoutLast
}
