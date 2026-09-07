package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//Задача: Реализация io.Writer с подсчётом байтов
//Условие:
//Вам необходимо создать тип CountingWriter, который реализует интерфейс io.Writer.
//Он должен оборачивать другой io.Writer и вести подсчёт общего количества записанных байтов.
//Структура CountingWriter должна содержать:
//поле writer (тип io.Writer), куда будет делегироваться запись;
//поле count (тип int64), хранящее суммарное количество переданных байтов.
//Реализуйте метод Write(p []byte) (n int, err error), который:
//вызывает метод Write у вложенного writer;
//увеличивает count на количество реально записанных байтов (возвращённое значение n);
//возвращает те же n и err, что и вложенный вызов.
//Напишите конструктор NewCountingWriter(w io.Writer) *CountingWriter, инициализирующий структуру.
//Напишите функцию TotalBytes(cw *CountingWriter) int64, возвращающую текущее значение счётчика.
//Продемонстрируйте работу:
//создайте CountingWriter, обёртывающий bytes.Buffer;
//запишите в него несколько строк (например, "Hello, " и "World!");
//выведите общее количество записанных байтов и содержимое буфера.
//Дополнительное задание (усложнение):
//Реализуйте аналогичный тип CountingReader для интерфейса io.Reader с методом Read,
//который также подсчитывает прочитанные байты.
//Продемонстрируйте чтение из strings.Reader через CountingReader.

type CountingWriter struct {
	writer io.Writer
	count  int64
}

func NewCountingWriter(writer io.Writer) *CountingWriter {
	return &CountingWriter{writer: writer}
}

func (cw *CountingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p)
	cw.count += int64(n)
	return n, err
}

func TotalBytes(cw *CountingWriter) int64 {
	return cw.count
}

type CountingReader struct {
	reader io.Reader
	count  int64
}

func NewCountingReader(reader io.Reader) *CountingReader {
	return &CountingReader{reader: reader}
}

func (cr *CountingReader) Read(p []byte) (n int, err error) {
	n, err = cr.reader.Read(p)
	fmt.Println("до", n, cr.count)
	cr.count += int64(n)
	fmt.Println("до", n, cr.count)
	return n, err
}

func TotalBytesRead(cr *CountingReader) int64 {
	return cr.count
}

func main() {
	buf := &bytes.Buffer{}
	cw := NewCountingWriter(buf)

	n1, _ := io.Copy(cw, bytes.NewBufferString("hello, "))
	n2, _ := io.Copy(cw, bytes.NewBufferString("world!"))

	fmt.Printf("Записано байт за раз: %d и %d\n", n1, n2)
	fmt.Printf("Всего записано байт: %d\n", TotalBytes(cw))
	fmt.Printf("Содержимое буфера: %q\n", buf.String())

	// Демонстрация CountingReader
	src := strings.NewReader("Go is awesome!")
	cr := NewCountingReader(src)

	data := make([]byte, 5)
	for {
		_, err := cr.Read(data)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("Всего прочитано байт: %d\n", TotalBytesRead(cr))

}
