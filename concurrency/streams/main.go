package main

import (
	"fmt"
	"io"
)

/* Custom implementation of io.Reader and io.Writer taken
from (see link below); detailed explanation of io.Streams in Go:
https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185
*/

type alphaReader struct {
	source  string
	current int
}

func newAlphaReader(source string) *alphaReader {
	return &alphaReader{source: source}
}

func alpha(b byte) byte {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return b
	}
	return 0
}

func (a *alphaReader) Read(b []byte) (int, error) {
	if a.current >= len(a.source) {
		return 0, io.EOF
	}
	x := len(a.source) - a.current
	n, bound := 0, 0
	if x >= len(b) {
		bound = len(b)
	} else if x <= len(b) {
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		char := alpha(a.source[a.current])
		if char != 0 {
			buf[n] = char
		}
		n++
		a.current++
	}
	copy(b, buf)
	return n, nil
}

type channelWriter struct {
	ch chan byte
}

func newChannelWriter() *channelWriter {
	return &channelWriter{make(chan byte, 1024)}
}

func (w *channelWriter) Write(b []byte) (int, error) {
	n := 0
	for _, v := range b {
		w.ch <- v
		n++
	}
	return n, nil
}

func (w *channelWriter) Channel() <-chan byte {
	return w.ch
}

func (w *channelWriter) Close() error {
	close(w.ch)
	return nil
}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()

	writer := newChannelWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me!"))
	}()
	for c := range writer.Channel() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
