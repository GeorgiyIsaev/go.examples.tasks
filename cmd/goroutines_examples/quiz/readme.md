Задача 1:

```go
// Что выведет?
func main() {
for i := 0; i < 3; i++ {
go func() {
fmt.Print(i, ", ")
}()
}
time.Sleep(500 * time.Millisecond)
```
Задача 2:
```go
// Что не так в коде?
func main() {
var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
```
Задача 3:
```go
// Что выведет?
func test1() {
ch := make(chan int)
ch <- 1
fmt.Println(<-ch)
//deadlock. Отправка в небуферизованный канал блокируется.
}
```

