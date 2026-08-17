package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} //len5 cap5
	s = s[1:4]                //2,3,4  len3 cap4
	s = s[:cap(s)]            //подвох cap(s) = 4
	//Значит расширим число действующих элементов до 4
	//значит теперь len4 cap4
	fmt.Println(s) //2,3,4,5
}
