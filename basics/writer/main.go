package main

import (
	"bytes"
	"fmt"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func RunWriter() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write to buffer")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())
}

type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{ch: make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func main() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me!"))
	}()
	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
