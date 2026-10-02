Задачи по синхронизации

1. Потокобезопасный счётчик на atomic
   Задача: реализовать тип Counter с методами Inc(), Add(delta int64), Value() int64. Мьютексы использовать нельзя.
   Тест: 100 горутин по 1000 Inc(), ожидаем 100000, запуск с -race.


2. Потокобезопасный кэш на RWMutex
   Задача: реализовать generic-кэш Cache[K comparable, V any] с методами Get, Set, Delete. Чтения должны идти параллельно, запись — эксклюзивно.
   Тест: параллельные Get и Set, запуск с -race.

3. ParallelMap на WaitGroup с сохранением порядка
   Задача: написать функцию ParallelMap(in []int, f func(int) int) []int, которая обрабатывает каждый элемент в отдельной горутине, но результат возвращает в исходном порядке.
   Тест: ParallelMap([]int{1,2,3}, func(x int) int { return x * x }) → [1,4,9], -race.

4. Worker pool с jobs/results и context
- Задача: реализовать функцию:

```
func WorkerPool(ctx context.Context, jobs <-chan int, workers int, f func(int) int) <-chan int
```
- Она запускает workers воркеров, читает jobs, пишет результаты в выходной канал и закрывает его после завершения всех воркеров. Должна уважать отмену ctx.
- Тест: закрыть jobs, прочитать всё из out, проверить отсутствие паники и утечек.












