package main

// Паттерн: фиксированное окно (fixed-size sliding window).
//
// Идея: размер окна k постоянен. Скользим окном по массиву:
// на каждом шаге убираем один элемент слева, добавляем один справа.
// Пересчёт "с нуля" не нужен — переиспользуем предыдущую сумму.
//
// Когда применять:
//   - "k последовательных элементов";
//   - "среднее по k элементам";
//   - "максимум/минимум в окне длины k";
//   - задачи про окна постоянного размера.
//
// Сложность: O(n) время, O(1) память.

// MaxSumSubarrayOfSizeK — максимальная сумма подмассива длины k.
//
// Пример: nums=[2,1,5,1,3,2], k=3 -> 9 (5+1+3).
func MaxSumSubarrayOfSizeK(nums []int, k int) int {
	if k <= 0 || len(nums) < k {
		return 0
	}

	// Первое окно: сумма первых k элементов.
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	// Сдвигаем окно: убираем nums[i-k], добавляем nums[i].
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

// FindAverages — средние значения всех окон длины k.
//
// Пример: nums=[1,3,2,6,-1,4,1,8,2], k=5
//
//	-> [2.2, 2.8, 2.4, 3.6, 2.8]
func FindAverages(nums []int, k int) []float64 {
	if k <= 0 || len(nums) < k {
		return []float64{}
	}
	result := make([]float64, 0, len(nums)-k+1)

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	result = append(result, float64(sum)/float64(k))

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		result = append(result, float64(sum)/float64(k))
	}
	return result
}

// MaxVowels — максимальное число гласных в подстроке длины k.
// Гласные: a, e, i, o, u (в нижнем регистре).
//
// Пример: s="abciiidef", k=3 -> 3 (подстрока "iii").
func MaxVowels(s string, k int) int {
	if k <= 0 || len(s) < k {
		return 0
	}

	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	maxCount := count

	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if isVowel(s[i-k]) {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
