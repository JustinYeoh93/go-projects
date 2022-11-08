package main

import "fmt"

type observers struct {
	id int
}

func (o *observers) OnNotify(e Event) {
	fmt.Printf("observer %d received event %d\n", o.id, e.Data)
}
