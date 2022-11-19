package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// reader := newAlphaReader("Hello! It's 8am, where is the sun?")
	// p := make([]byte, 4)

	// for {
	// 	n, err := reader.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(string(p[:n]))
	// }
	// fmt.Println()

	reader := newChainReader(strings.NewReader("Hello!, its 9am, where the sun"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
