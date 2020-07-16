package test

import (
	"fmt"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	num := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}
