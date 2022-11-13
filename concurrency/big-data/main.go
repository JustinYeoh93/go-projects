package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", concurrent_processing("data/test.txt", 3, 20))
}
