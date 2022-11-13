package mod1

import (
	"fmt"
)

type User struct {
	Name string
}

func (u *User) SayHi() string {
	return fmt.Sprintf("Hi there, I'm %s.", u.Name)
}

func (u *User) SayBye() string {
	return "Bye!"
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}
