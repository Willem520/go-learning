package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	fmt.Println("======start")
	tes1()
	fmt.Println("======end")
}

func tes1() {
	fmt.Println("======start test1")
	defer fmt.Println("======defer in test1")
	fmt.Println("======end test1")
}

func TestPointer(t *testing.T) {
	i, j := 12, 22
	p := &i         //指向i
	fmt.Println(*p) //通过指针获取i的值
	*p = 21         //通过指针设置i的值
	fmt.Println(i)

	p = &j       //指向j
	*p = *p / 11 //通过指针对j进行除法运算
	fmt.Println(j)
}
