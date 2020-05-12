package test

import (
	"fmt"
	"testing"
)

/**
多态的展示
*/

//定义一个通知类行为的接口
type notifier interface {
	notify()
}

type usr struct {
	name  string
	email string
}

func (u *usr) notify() {
	fmt.Printf("sending usr email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func TestNotify(t *testing.T) {
	//创建一个usr类型的值，并传给sendNotification
	u := usr{"Bill", "bill@email.com"}
	sendNotification(&u)

	//创建一个admin类型的值，并传给sendNotification
	a := admin{"Lisa", "lisa@email.com"}
	sendNotification(&a)
}
