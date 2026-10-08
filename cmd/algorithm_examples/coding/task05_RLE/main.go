package main

import (
	"fmt"
	"strconv"
	"strings"
)

//5. RLE-сжатие строки
//Условие: Дана строка из строчных латинских букв. Вернуть сжатие: символ + количество подряд идущих. Количество писать всегда.
//Пример:
//"aaabbc" -> "a3b2c1"
//"abcd" -> "a1b1c1d1"
//Сигнатура:
//func Compress(s string) string

func Compress(s string) string {
	if s == "" {
		return ""
	}
	var sb strings.Builder
	runes := []rune(s)

	current := runes[0]
	count := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == current {
			count++
		} else {
			// серия закончилась — записываем её
			sb.WriteRune(current)
			sb.WriteString(strconv.Itoa(count))
			// начинаем новую серию
			current = runes[i]
			count = 1
		}
	}

	// не забываем последнюю серию — она "висит" в current/count
	sb.WriteRune(current)
	sb.WriteString(strconv.Itoa(count))

	return sb.String()
}

func main() {
	fmt.Println(Compress("aaabbc")) // a3b2c1
	fmt.Println(Compress("abcd"))   // a1b1c1d1
	fmt.Println(Compress("aabbaa")) // a2b2a2
	fmt.Println(Compress(""))       // ""
	fmt.Println(Compress("aaaa"))   // a4
}
