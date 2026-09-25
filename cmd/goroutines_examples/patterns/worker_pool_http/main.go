package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Click — то, что отправляем на ручку.
type Click struct {
	AuthorID int   `json:"author_id"`
	UserID   int   `json:"user_id"`
	Ts       int64 `json:"ts"`
}

// Result — результат отправки одного клика.
type Result struct {
	Click Click
	Err   error
}

// sendClick делает один HTTP-запрос. Это «единица работы» воркера.
func sendClick(ctx context.Context, client *http.Client, baseURL string, c Click) error {
	body, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/api/clicks", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	// 5xx и 429 — обычно стоит ретраить, 4xx — нет.
	if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests {
		return fmt.Errorf("server error: status %d", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("client error: status %d", resp.StatusCode)
	}
	return nil
}

// worker читает клики из jobs, шлёт их и пишет результат в results.
func worker(
	ctx context.Context,
	id int,
	client *http.Client,
	baseURL string,
	jobs <-chan Click,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// Контекст отменили — выходим, не дожидаясь новых задач.
			return
		case click, ok := <-jobs:
			if !ok {
				return // канал закрыт, задач больше нет
			}
			err := sendClick(ctx, client, baseURL, click)
			// Результат тоже отправляем с учётом ctx,
			// чтобы не зависнуть, если main перестал читать.
			select {
			case results <- Result{Click: click, Err: err}:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	const (
		numWorkers = 5
		numClicks  = 100
		baseURL    = "http://localhost:8080" // замени на свой адрес
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Один общий HTTP-клиент на всех воркеров.
	// У него внутри свой пул соединений — переиспользуем TCP.
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        numWorkers * 2,
			MaxIdleConnsPerHost: numWorkers,
			IdleConnTimeout:     30 * time.Second,
		},
	}

	jobs := make(chan Click)
	results := make(chan Result)

	var wg sync.WaitGroup

	// 1. Стартуем воркеров.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, client, baseURL, jobs, results, &wg)
	}

	// 2. Отдельная горутина кладёт клики в jobs и закрывает канал.
	go func() {
		defer close(jobs)
		for i := 1; i <= numClicks; i++ {
			click := Click{
				AuthorID: rand.Intn(10) + 1,
				UserID:   i,
				Ts:       time.Now().Unix(),
			}
			select {
			case jobs <- click:
			case <-ctx.Done():
				return // не успели отправить — выходим
			}
		}
	}()

	// 3. Горутина, которая после всех воркеров закроет results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Собираем результаты.
	var ok, failed int
	for res := range results {
		if res.Err != nil {
			failed++
			fmt.Printf("click user=%d FAILED: %v\n", res.Click.UserID, res.Err)
			continue
		}
		ok++
	}

	fmt.Printf("done: ok=%d failed=%d\n", ok, failed)

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		fmt.Println("warning: stopped by timeout, not all clicks were sent")
	}
}
