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
