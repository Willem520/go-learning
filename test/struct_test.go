package test

import (
	"fmt"
	"testing"
)

type Vertex struct {
	X, Y int
}

func TestStruct(t *testing.T) {
	v1 := Vertex{1, 2} //创建一个Vertex类型的结构体
	v2 := Vertex{X: 1} //Y:0被隐式创建
	v3 := Vertex{}     //X:0 Y:0
	p := &Vertex{1, 2} //创建一个*Vertex类型的结构体（指针）
	fmt.Println(v1, p, v2, v3)
}
