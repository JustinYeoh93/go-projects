package mod1

import "testing"

func Test_User_SayHi(t *testing.T) {
	name := "James"
	expected := "Hi there, I'm James."

	u := NewUser(name)
	actual := u.SayHi()

	if actual != expected {
		t.Errorf("execpted '%s' but got '%s'", expected, actual)
	}
}

func Test_User_SayBye(t *testing.T) {
	name := "James"
	expected := "Bye!"

	u := NewUser(name)
	actual := u.SayBye()

	if actual != expected {
		t.Errorf("execpted '%s' but got '%s'", expected, actual)
	}
}
