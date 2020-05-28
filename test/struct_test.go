package test

import (
	"fmt"
	"testing"
)

type vertex struct {
	X, Y int
}

func TestStruct(t *testing.T) {
	v1 := vertex{1, 2} //创建一个Vertex类型的结构体
	v2 := vertex{
		X: 2,
		Y: 3,
	}
	v3 := vertex{X: 1} //Y:0被隐式创建
	v4 := vertex{}     //X:0 Y:0
	p := &vertex{1, 2} //创建一个*Vertex类型的结构体（指针）
	fmt.Println(v1, v2, v3, v4, p)
}
