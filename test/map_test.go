package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("the value:", m["answer"])

	m["answer"] = 48
	fmt.Println("the value:", m["answer"])

	delete(m, "answer")
	fmt.Println("the value:", m["answer"])

	v, ok := m["answer"]
	fmt.Println("the value", v, "present?", ok)

	color := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range color {
		fmt.Printf("Key: %s Value:%s\n", key, value)
	}
}
