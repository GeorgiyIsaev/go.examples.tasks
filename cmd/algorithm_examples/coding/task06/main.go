package main

import "fmt"

//Проверка правильности скобок
//   Условие: Дана строка, состоящая только из символов (, ), {, }, [, ]. Вернуть true, если скобки расставлены корректно: каждая открывающая закрыта соответствующей закрывающей, и порядок вложенности соблюдён.
//    Сигнатура:
//   func IsValid(s string) bool

func IsValid(s string) bool {
	if s == "" {
		return true
	}

	b := []byte(s)
	/// (([{}]))
	// (
	//  (
	//   [
	//    {
	//     }
	//      ]
	//		 )
	//  	  )
	sl := make([]byte, 0)

	for i := 0; i < len(b); i++ {

		switch b[i] {
		case '(':
			sl = append(sl, '(')
		case '[':
			sl = append(sl, '[')
		case '{':
			sl = append(sl, '{')

		case ')':
			if len(sl) == 0 || sl[len(sl)-1] != byte('(') {
				return false
			}
			sl = sl[:len(sl)-1]
		case ']':
			if len(sl) == 0 || sl[len(sl)-1] != byte('[') {
				return false
			}
			sl = sl[:len(sl)-1]
		case '}':
			if len(sl) == 0 || sl[len(sl)-1] != byte('{') {
				sl = sl[:len(sl)-1]
			}
			sl = sl[:len(sl)-1]
		}
	}
	return len(sl) == 0
}

func main() {
	str := "(, ), {, }, [, ]"
	fmt.Println(str, IsValid(str))

	str = "([{}])"
	fmt.Println(str, IsValid(str))

	str = "(([{}]))"
	fmt.Println(str, IsValid(str))

	str = "([{(}])]" //это false
	fmt.Println(str, IsValid(str))

	str = "(([)])" //это false
	fmt.Println(str, IsValid(str))

	str = "()])" //это false
	fmt.Println(str, IsValid(str))

	fmt.Println(IsValid("("))
	fmt.Println(IsValid("([{}]"))

}
