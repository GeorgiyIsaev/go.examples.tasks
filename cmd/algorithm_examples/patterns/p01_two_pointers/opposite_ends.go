package main

// Паттерн: два указателя навстречу друг другу (opposite ends).
//
// Идея: left указывает на начало, right на конец. На каждом шаге
// либо left++, либо right--, в зависимости от условия. Указатели
// движутся навстречу и встречаются в середине.
//
// Когда применять:
//   - массив отсортирован;
//   - нужно обработать пару/симметрию;
//   - in-place без дополнительной памяти.
//
// Сложность: O(n) время, O(1) память.

// ReverseString разворачивает срез байтов in-place.
// Классический пример сближающихся указателей.
func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// IsPalindrome проверяет, является ли строка палиндромом.
// Учитываются только буквы и цифры, регистр не важен.
// Русские буквы не поддерживаются — работаем с ASCII для простоты.
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlnum(s[left]) {
			left++
		}
		for left < right && !isAlnum(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlnum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// TwoSumSorted ищет два индекса в отсортированном массиве,
// сумма элементов по которым равна target. Возвращает nil,
// если такой пары нет.
//
// Это ядро всех задач «Two Sum» на отсортированных данных
// и строительный блок для Three Sum / Four Sum.
func TwoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		switch {
		case sum == target:
			return []int{left, right}
		case sum < target:
			left++
		default:
			right--
		}
	}
	return nil
}
