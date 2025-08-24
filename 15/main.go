package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return strings.Repeat("f", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	if len(runes) < 100 {
		justString = v
		return
	}
	result := make([]rune, 100)
    copy(result, runes[:100])
	justString = string(result)
}

func main() {
	someFunc()
	fmt.Printf("длина justString: %d символов\n", len(justString))
}
