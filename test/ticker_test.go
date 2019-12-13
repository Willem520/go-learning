package test

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 5)
	wait := sync.WaitGroup{}
	wait.Add(1)
	loop := 0
	go func() {
		for _ = range ticker.C {
			if loop < 3 {
				log.Printf("======loop:" + strconv.Itoa(loop))
				loop++
				continue
			}
			wait.Done()
		}
	}()
	wait.Wait()
	fmt.Println("======")
}
