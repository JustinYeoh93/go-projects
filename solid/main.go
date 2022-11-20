package main

import (
	"fmt"
	"solid/ocp"
	"solid/srp"
)

// base entity type
type vehicle struct {
	name string
}

func (v vehicle) getName() string {
	return v.name
}

// subtype motorcycle that is composed with vehicle
type motorcycle struct {
	vehicle
	wheel int
}

func main() {
	// SRP
	var s []srp.Shape
	s = append(s, srp.NewCircle(5))
	s = append(s, srp.NewSquare(2))

	for _, sh := range s {
		fmt.Println(sh.GetArea())
	}

	// OCP
	var c ocp.Calculator
	sum := c.SumAreas(&ocp.Circle{Radius: 5}, &ocp.Square{SideLen: 5}, &ocp.EquilateralTriangle{Base: 2, Height: 3})
	fmt.Println(sum)

	m := motorcycle{vehicle{"potato"}, 2}
	fmt.Println(m.getName())
	fmt.Println(m.name)
}
