package mod1

import "testing"

func Test_User(t *testing.T) {
	name := "James"
	expected := "Hi there, I'm James."
	u := NewUser(name)

	if u.SayHi() != expected {
		t.Errorf("execpted '%s' but got '%s'", expected, u.SayHi())
	}
}
