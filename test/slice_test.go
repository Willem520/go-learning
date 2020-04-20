package test

import (
	"fmt"
	"testing"
)

/*
对于容量为k的切片slince[i:j]来说，长度：j-i，容量：k-i
*/
func TestSlice(t *testing.T) {
	//创建一个字符串切片，长度和容量都是5
	s1 := make([]string, 5)
	fmt.Println(s1)

	//创建一个整形切片，其长度为3，容量为5
	s2 := make([]int, 3, 5)
	fmt.Println(s2)

	//创建长度和容量为5的切片
	primes := []int{2, 3, 5, 7, 11, 13}
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
