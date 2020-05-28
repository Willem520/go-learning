package test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var a [5]string
	a[0] = "Hello"
	a[2] = "Go"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//使用...根据初始化时的大小确认容量
	auto := [...]int{1, 2, 3}
	fmt.Println(auto)

	//初始化数组指定下标的值
	arr := [5]int{1: 21, 3: 10}
	fmt.Println(arr)
}
