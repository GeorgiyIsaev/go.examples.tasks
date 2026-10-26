package main

// Две классики DP по двум строкам: LCS и расстояние Левенштейна.
//
// Общая схема: dp[i][j] = оптимум для ПРЕФИКСОВ s[:i] и t[:j].
// Ключевой вопрос — что делать с последними символами:
//   - они совпали? → dp[i-1][j-1] + 1 (берём «пару»);
//   - нет? → максимум/минимум из трёх вариантов «отбросить что-то».
// Ответ — dp[len(s)][len(t)] (правый нижний угол таблицы).
// Сложность: O(|s|·|t|) время и память.

// LongestCommonSubsequence — длина ЛКС двух строк (LeetCode 1143).
// Подпоследовательность, не обязательно подстрока.
//
// Пример: "abcde" / "ace" -> 3 ("ace").
func LongestCommonSubsequence(s, t string) int {
	m, n := len(s), len(t)
	// dp[i][j] — ЛКС префиксов s[:i] и t[:j].
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 // символы совпали
			} else {
				// отбросили символ из s или из t — берём лучшее
				a, b := dp[i-1][j], dp[i][j-1]
				if a > b {
					dp[i][j] = a
				} else {
					dp[i][j] = b
				}
			}
		}
	}
	return dp[m][n]
}

// MinDistance — минимальное число операций, чтобы превратить word1
// в word2 (расстояние Левенштейна, LeetCode 72).
// Операции: вставить, удалить, заменить — каждая стоит 1.
//
// dp[i][j] — сколько операций, чтобы превратить word1[:i] в word2[:j]:
//   - совпали → dp[i-1][j-1], операция не нужна;
//   - замена  → dp[i-1][j-1] + 1;
//   - вставка → dp[i][j-1] + 1   (вставили word2[j-1] в word1);
//   - удаление→ dp[i-1][j] + 1   (удалили word1[i-1]).
//
// Пример: "horse" -> "ros" -> 3; "intention" -> "execution" -> 5.
func MinDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// База: из пустой строки — только вставки/удаления.
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			best := dp[i-1][j-1] + 1 // замена последних символов
			if dp[i][j-1]+1 < best {
				best = dp[i][j-1] + 1 // вставка
			}
			if dp[i-1][j]+1 < best {
				best = dp[i-1][j] + 1 // удаление
			}
			dp[i][j] = best
		}
	}
	return dp[m][n]
}
