package main

import (
	"fmt"
	"math/rand"
)

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
)

func main() {
	const Purple Color = "purple"

	arr := []Color{Red, Red, Blue, Blue, Red, Purple}

	for i := 0; i < 10; i++ {
		var colorNow = arr[rand.Intn(len(arr))]
		fmt.Printf("Color now is: %s\n", colorNow)
		fmt.Printf("Color is red: %t\n", isRed(colorNow))
	}
}

func isRed(color Color) bool {
	return color == Red
}
