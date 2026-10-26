package main

// 0/1 Рюкзак — «матерь всех DP-задач на выбор».
//
// Дано: предметы с весом w[i] и ценностью v[i], вместимость W.
// Каждый предмет — взять целиком или не брать (0/1), максимум ценности.
//
// Состояние: dp[i][c] — лучшая ценность, используя первые i предметов
// и вместимость не больше c. Переход для i-го предмета:
//
//	dp[i][c] = max(dp[i-1][c],              // не берём
//	               dp[i-1][c-w[i]] + v[i])  // берём (если влезает)
//
// Память: dp[i] зависит только от ряда i-1 => один ряд + обход
// справа налево (!) — так «берём» ещё не переиспользует обновлённые
// значения текущего ряда. Обход влево дал бы «взять много раз»
// (это уже unbounded knapsack).
//
// Сложность: O(n·W) время, O(W) память. Рюкзак — pseudo-polynomial:
// по «длинным» числам W экспоненциально растёт.

// Knapsack01 — максимум ценности при вместимости W (классический
// 0/1 knapsack, встречается под номером вроде LC416/474 в адаптированных
// формулировках — здесь базовая версия без номера).
//
// Пример: W=7, w=[1,3,4,5], v=[1,4,5,7] -> 9 (предметы 3+4).
func Knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1) // dp[c] — лучшая ценность при ёмкости c
	for i := 0; i < len(weights); i++ {
		w, v := weights[i], values[i]
		// Справа налево: каждым предметом пользуемся один раз.
		for c := capacity; c >= w; c-- {
			take := dp[c-w] + v
			if take > dp[c] {
				dp[c] = take
			}
		}
	}
	return dp[capacity]
}

// SubsetSum — можно ли набрать ровно target суммой подмножества
// (LeetCode 416 Partition Equal Subset Sum в формулировке «ровно
// target» — тот же рюкзак с ценностью = весу).
//
// Идея: та же таблица, но boolean: dp[c] = «сумма c достижима».
// Старт: dp[0] = true; на каждом числе обновляем справа налево.
//
// Пример: [1,5,11,5], target=11 -> true (5+6? нет: 5+5+1 = 11).
func SubsetSum(nums []int, target int) bool {
	dp := make([]bool, target+1)
	dp[0] = true
	for _, x := range nums {
		for c := target; c >= x; c-- {
			if dp[c-x] {
				dp[c] = true
			}
		}
	}
	return dp[target]
}

// PartitionEqual — можно ли разбить массив на два равных подмножества
// (LeetCode 416): рюкзак на target = сумма/2.
//
// Пример: [1,5,11,5] -> true (1+5+5 = 11).
func PartitionEqual(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 != 0 {
		return false
	}
	return SubsetSum(nums, total/2)
}
