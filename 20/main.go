package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		fmt.Println(reverseStr(sc.Text()))
	}
}

func reverseStr(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	fmt.Println(runes)
	reverseRunes(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}
	if start < len(runes) {
		reverseRunes(runes, start, len(runes)-1)
	}

	return string(runes)
}

func reverseRunes(runes []rune, x, y int) {
	for x < y {
		runes[x], runes[y] = runes[y], runes[x]
		x++
		y--
	}
}
