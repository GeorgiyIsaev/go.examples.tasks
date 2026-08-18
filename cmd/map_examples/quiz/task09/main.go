package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[int]int{}
	go func() {
		for i := 0; i < 1000; i++ {
			m[i] = i
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			m[i] = i * 2
		}
	}()
	//две горутины без мьютексов
	//результат гонка, возможно паника
	//нужно sync.Mutex или sync.Map
	time.Sleep(2 * time.Second)
	fmt.Println(len(m)) //размер 1000
	fmt.Println(m)      //результат не гарантирован
}
