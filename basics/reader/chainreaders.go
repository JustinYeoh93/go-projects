package main

import (
	"io"
)

type chainReader struct {
	reader io.Reader
}

func newChainReader(reader io.Reader) *chainReader {
	return &chainReader{reader}
}

func (a *chainReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}
