package main

import (
	"fmt"

	"github.com/JustinYeoh93/go-projects/semver/mod1"
)

func main() {
	u := mod1.NewUser("test")
	fmt.Println(u.SayHi())
	fmt.Println(u.SayBye())
}
