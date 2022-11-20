package srp

import "math"

type Shape interface {
	GetArea() float32
}

type circle struct {
	radius float32
}

func NewCircle(r float32) Shape {
	return circle{r}
}

func (c circle) GetArea() float32 {
	// // breaking srp because the function is calculating and printing the result
	// fmt.Println(c.radius * c.radius * math.Pi)
	return c.radius * c.radius * math.Pi
}

type square struct {
	sideLen float32
}

func NewSquare(s float32) Shape {
	return square{s}
}

func (s square) GetArea() float32 {
	// // breaking srp because the function is calculating and printing the result
	// fmt.Println(s.sideLen * s.sideLen)
	return s.sideLen * s.sideLen
}
