package main

import "fmt"

//Дана map map[string]interface{},
//которая может содержать значения разных типов:
//int, string, float64, []interface{}.
//Напишите функцию extractInts(m map[string]interface{}) []int,
//которая обходит map рекурсивно
//и собирает все целочисленные значения (типа int и int64)
//в плоский слайс.

func extractInts(m map[string]interface{}) []int {
	var result []int

	var recurse func(interface{})
	recurse = func(v interface{}) {
		switch v := v.(type) {
		case map[string]interface{}:
			for _, subVal := range v {
				recurse(subVal)
			}
		case []interface{}:
			for _, elem := range v {
				recurse(elem)
			}
		case int:
			result = append(result, v)
		case int64:
			result = append(result, int(v))
		default:
		}
	}

	recurse(m)
	return result
}

func main() {
	data := map[string]interface{}{
		"a": 10,
		"b": "hello",
		"c": map[string]interface{}{
			"d": 20,
			"e": []interface{}{30, "x", 40},
		},
	}
	result := extractInts(data) // [10, 20, 30, 40]
	fmt.Println(result)
}
