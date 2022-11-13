package main

import (
	"fmt"

	"github.com/JustinYeoh93/private-go-modules/user"
)

func main() {
	u := user.NewUser("test")
	fmt.Println(u.GetName())
}
