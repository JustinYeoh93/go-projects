package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	wrapperPanic()
	fmt.Println("After wrapperPanic()")
}

func wrapperPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()

	mayPanic()
}
