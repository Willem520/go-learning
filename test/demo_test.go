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
