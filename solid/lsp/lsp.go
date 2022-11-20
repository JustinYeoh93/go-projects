package lsp

import "fmt"

type transport interface {
	getName() string
}

// base entity type
type vehicle struct {
	name string
}

func (v vehicle) getName() string {
	return v.name
}

// subtype car that is composed with vehicle
type car struct {
	vehicle
	name  string
	wheel int
	gates int
}

// subtype motorcycle that is composed with vehicle
type motorcycle struct {
	vehicle
	name  string
	wheel int
}

type printer struct{}

func (p printer) printTransportName(t transport) {
	fmt.Println(t.getName())
}
