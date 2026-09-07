package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//Интерфейс для логирования (несколько реализаций)
//Определите интерфейс Logger с методами:
//Log(message string)
//Logf(format string, args ...interface{}) (форматированный вывод).
//Реализуйте два типа:
//ConsoleLogger – пишет сообщения в os.Stdout (с меткой времени, можно использовать log пакет);
//FileLogger – пишет сообщения в файл (путь передаётся при создании).
//Напишите функцию LogMessages(logger Logger, messages []string),
//которая принимает логгер и список сообщений, и логирует каждое из них с помощью Log.
//Также продемонстрируйте использование Logf.
//Проверьте работу с обоими реализациями.

type Logger interface {
	Log(message string)
	Logf(format string, args ...interface{})
}

type ConsoleLogger struct {
}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), message)
}
func (c ConsoleLogger) Logf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	c.Log(message)
}

type FileLogger struct {
	file *os.File
}

func (f FileLogger) Log(message string) {
	log.SetOutput(f.file)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), message)
}
func (f FileLogger) Logf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	f.Log(message)
}

func NewFileLogger(filename string) (*FileLogger, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: f}, nil
}

func main() {
	// Консольный логгер
	console := ConsoleLogger{}
	console.Log("Start")
	console.Log("Processing")
	console.Log("End")
	console.Logf("User %s logged in", "Alice")

	// Файловый логгер
	fileLogger, _ := NewFileLogger("app.log")
	fileLogger.Log("Error occurred")
	fileLogger.Log("Recovery")
}
