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
Задча 4:
```go
// Вопрос: всегда ли выведется 1000?
func ex1() {
var counter int
var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
			//!!! data race. counter++ — это не атомарная операция
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
```
Задач 5:
```go
//Вопрос: что не так?
func ex(ctx context.Context) {
ch := make(chan int)

	go func() {
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		return
	case v := <-ch:
		fmt.Println(v)
	}
}
```

Задач 6:
```go
// Что не так с использованием WaitGroup
func ex() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done() // counter = 0, Wait может проснуться
		wg.Add(1) // новая партия
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Println("second done")
		}()
	}()
	//Нарушен порядок
	wg.Wait()
	fmt.Println("main done")
}
```

Задач 7:
```go

```

Задач 8:
```go

```
