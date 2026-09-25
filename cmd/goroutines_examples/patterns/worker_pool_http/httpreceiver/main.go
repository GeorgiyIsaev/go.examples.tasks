package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type Click struct {
	AuthorID int   `json:"author_id"`
	UserID   int   `json:"user_id"`
	Ts       int64 `json:"ts"`
}

var (
	received atomic.Int64 // сколько всего кликов приняли
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/clicks", handleClick)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           logMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Запускаем сервер в отдельной горутине, чтобы main мог ждать сигнал.
	go func() {
		log.Println("listening on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()

	// Ждём Ctrl+C или SIGTERM.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Printf("bye. total clicks received: %d", received.Load())
}

func handleClick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var c Click
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // ловим опечатки в полях
	if err := dec.Decode(&c); err != nil {
		log.Printf("bad request: %v", err)
		http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
		return
	}

	if c.AuthorID == 0 || c.UserID == 0 {
		log.Printf("validation failed: %+v", c)
		http.Error(w, "author_id and user_id required", http.StatusBadRequest)
		return
	}

	n := received.Add(1)
	log.Printf("#%d  author=%d user=%d ts=%d  remote=%s",
		n, c.AuthorID, c.UserID, c.Ts, r.RemoteAddr)

	// Можно специально возвращать ошибку для тестов ретраев.
	// Раскомментируй, если хочешь проверить, как клиент обрабатывает 500.
	// if n%5 == 0 {
	// 	http.Error(w, "simulated failure", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte(`{"status":"accepted"}`))
}

// logMiddleware логирует метод, путь и код ответа.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("HTTP %s %s -> %d (%s)", r.Method, r.URL.Path, rw.status, time.Since(start))
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
