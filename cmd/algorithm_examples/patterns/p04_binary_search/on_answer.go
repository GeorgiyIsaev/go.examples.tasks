package main

// Паттерн: бинарный поиск по ответу (binary search on answer).
//
// Самый мощный и самый недооценённый подвид. Идея:
//   - ответ лежит в диапазоне [low, high];
//   - есть монотонная функция feasible(x): «подходит ли ответ x»;
//   - если feasible(x) истинно, то истинно и для всех больших (или
//     меньших — зависит от задачи) x;
//   - значит, можно бинарным поиском найти границу.
//
// Мы НЕ ищем элемент в массиве. Мы ищем число — ответ задачи.
//
// Шаблон:
//   low, high := минимально возможный ответ, максимально возможный
//   for low < high {
//       mid := low + (high-low)/2
//       if feasible(mid) { high = mid }  // или low = mid + 1 — зависит
//       else { low = mid + 1 }
//   }
//   return low

// MySqrt — целочисленный квадратный корень (floor).
//
// Пример: 4 -> 2. 8 -> 2. 0 -> 0. 1 -> 1.
//
// Используем деление вместо умножения — защита от overflow при
// больших x (mid*mid может переполнить int).
func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x/2
	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		}
		if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

// MinEatingSpeed — минимальная скорость, с которой Коко съест все
// бананы за h часов. piles[i] — кучка с piles[i] бананами.
// За один час Коко ест не более speed бананов и переходит к
// следующей кучке только когда текущая закончилась.
//
// Пример: piles=[3,6,7,11], h=8 -> 4.
//
// Идея: ответ в [1, max(piles)]. Проверяем, можно ли управиться
// со скоростью speed за h часов. Функция canFinish монотонна:
// чем больше speed, тем меньше часов.
func MinEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	canFinish := func(speed int) bool {
		hours := 0
		for _, p := range piles {
			hours += (p + speed - 1) / speed // ceil(p/speed)
		}
		return hours <= h
	}

	left, right := 1, maxPile
	for left < right {
		mid := left + (right-left)/2
		if canFinish(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// ShipWithinDays — минимальная грузоподъёмность корабля, чтобы
// перевезти все грузы weights за days дней. Корабль ходит раз в
// день и за один день берёт любую подпоследовательность подряд
// идущих грузов, суммарно не превышающую capacity.
//
// Пример: weights=[1,2,3,4,5,6,7,8,9,10], days=5 -> 15.
//
// Нижняя граница: max(weights) — иначе самый тяжёлый не влезет.
// Верхняя: сумма всех — можно перевезти за один день.
func ShipWithinDays(weights []int, days int) int {
	maxWeight, totalWeight := 0, 0
	for _, w := range weights {
		if w > maxWeight {
			maxWeight = w
		}
		totalWeight += w
	}

	canShip := func(capacity int) bool {
		needDays := 1
		curLoad := 0
		for _, w := range weights {
			if curLoad+w > capacity {
				needDays++
				curLoad = 0
			}
			curLoad += w
		}
		return needDays <= days
	}

	left, right := maxWeight, totalWeight
	for left < right {
		mid := left + (right-left)/2
		if canShip(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
