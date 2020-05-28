package test

import (
	"fmt"
	"testing"
)

type user struct {
	name  string
	email string
}

//使用值接收者实现
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func TestFunc(t *testing.T) {
	a := user{"jack", "jack@email.com"}
	a.notify()

	b := &user{"rose", "rose@email.com"}
	b.notify()

	a.changeEmail("jack@newemail.com")
	a.notify()

	b.changeEmail("rose@newemail.com")
	b.notify()
}
