package test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Go"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func TestSlice(t *testing.T) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])

	//截取切片使其长度为0
	s := primes[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//扩展其长度
	s = s[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//舍弃前两个值
	s = s[2:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//更改切片的元素会修改其底层数组中的对应元素
	name := [4]string{"test1", "test2", "test3", "test4"}
	a := name[0:2]
	b := name[1:3]
	fmt.Println(a, b)

	b[0] = "test0"
	fmt.Println(a, b)
	fmt.Println(name)
}
