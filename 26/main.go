package main

import (
	"fmt"
	"unicode"
)

func isUnique(str string) bool {
	m := make(map[rune]struct{})

	for _, v := range str {
		v1 := unicode.ToLower(v)
		if _, ok := m[v1]; ok {
			return false
		}
		m[v1] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
	fmt.Println(isUnique("akfugs"))
}
