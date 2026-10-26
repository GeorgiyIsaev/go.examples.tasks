package main

// Паттерн: монотонный стек.
//
// Идея: в стеке держим индексы элементов, для которых ответ ещё не
// найден, а сами значения — МОНОТОННЫЕ (убывающие — если ищем
// «следующий больший», возрастающие — если «следующий меньший»).
//
// Проходя справа налево/слева направо, мы «закрываем» отложенных
// кандидатов: пока вершина стека хуже текущего элемента — текущий
// и есть ответ для вершины, pop. Каждый элемент входит и выходит
// ровно один раз => O(n).
//
// Сложность: O(n) время, O(n) память (в худшем — весь вход в стеке,
// например на строго убывающей последовательности).
//
// Общая форма (слева направо, «следующий больший»):
//
//	for i := 0; i < n; i++ {
//	    for стек не пуст && nums[i] > nums[стек.верх] {
//	        ответ[стек.верх] = nums[i]
//	        pop
//	    }
//	    push(i)
//	}

// NextGreaterElement — для каждого элемента из nums1 найти ближайший
// больший справа в nums2 (nums1 — подмножество nums2, LeetCode 496).
// Нет такого — -1.
//
// Пример: nums1=[4,1,2], nums2=[1,3,4,2] -> [-1,3,-1].
// (У 4 справа ничего больше; у 1 — 3; у 2 — ничего.)
func NextGreaterElement(nums1, nums2 []int) []int {
	// Ответ строим по всему nums2, потом вырезаем нужные позиции.
	index := map[int]int{} // значение -> индекс в nums2
	for i, v := range nums2 {
		index[v] = i
	}

	full := nextGreater(nums2)
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = full[index[v]]
	}
	return result
}

// nextGreater — базовый проход «вправо к ближайшему большему»
// по всему массиву. Возвращает -1 там, где большего нет.
func nextGreater(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := range answer {
		answer[i] = -1
	}

	stack := []int{} // индексы, ответ ещё не найден (убывающие значения)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = nums[i]
		}
		stack = append(stack, i)
	}
	return answer
}

// NextGreaterElements — то же, но массив «зациклен»: последний элемент
// считается идущим перед первым (LeetCode 503).
//
// Трюк: проходим по индексам 0..2n-1, индекс в массиве — i % n.
// В стек держим индексы «первого круга», но сравниваем значения по
// модулю n. Хватает одного прохода в 2n.
//
// Пример: [1,2,1] -> [2,-1,2].
func NextGreaterElements(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := range answer {
		answer[i] = -1
	}

	stack := []int{} // реальные индексы 0..n-1
	for i := 0; i < 2*n; i++ {
		cur := nums[i%n]
		for len(stack) > 0 && cur > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = cur
		}
		// Кладём только первый круг — второй лишь «закрывает» хвост.
		if i < n {
			stack = append(stack, i)
		}
	}
	return answer
}
