package ocp

import "math"

type shape interface {
	area() float32
}

type Circle struct {
	Radius float32
}

func (c *Circle) area() float32 {
	return c.Radius * c.Radius * math.Pi
}

type Square struct {
	SideLen float32
}

func (s *Square) area() float32 {
	return s.SideLen * s.SideLen
}

type EquilateralTriangle struct {
	Base   float32
	Height float32
}

func (e *EquilateralTriangle) area() float32 {
	return e.Base * e.Height / 2
}

type Calculator struct{}

// // This breaks Open/ Close principal because Calculator has a method SumAreas
// // that defines as a parameter "shapes" of interface{} type. This is breaking the
// // principle because with the introduction of shape triangle, we will need to modify
// // SumAreas switch section
// func (c Calculator) SumAreas(shapes ...interface{}) float32 {
// 	var sum float32

// 	for _, shape := range shapes {
// 		switch shape.(type) {
// 		case Circle:
// 			r := shape.(Circle).radius
// 			sum += r * r * math.Pi
// 		case Square:
// 			s := shape.(Square).sideLen
// 			sum += s * s
// 		}
// 	}

// 	return sum
// }

func (c *Calculator) SumAreas(shapes ...shape) float32 {
	var sum float32
	for _, s := range shapes {
		sum += s.area()
	}

	return sum
}
