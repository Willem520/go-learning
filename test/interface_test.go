package test

import (
	"fmt"
	"math"
	"testing"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vert struct {
	X, Y float64
}

func (v *Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test(t *testing.T) {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vert{3, 4}

	a = f  // a MyFloat实现了Abser
	a = &v //a *Vert实现了Abser

	fmt.Println(a.Abs())
}

/*
类型断言
*/
func Test2(t *testing.T) {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //报错panic
	fmt.Println(f)
}

/*
类型选择
*/
func Test3(t *testing.T) {
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
