package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go func() {
		for {
			fmt.Println("hello")
			time.Sleep(time.Millisecond * 700)
		}
	}()
	time.Sleep(time.Second * 5)
}

func TestChannel(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
