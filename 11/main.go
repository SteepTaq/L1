package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Printf("A= %v\n", A)
	fmt.Printf("B = %v\n", B)

	intersection := findIntersection(A, B)
	fmt.Printf("Пересечение = %v\n", intersection)
}

func findIntersection(set1, set2 []int) []int {
	m := make(map[int]struct{})

	for _, elem := range set1 {
		m[elem] = struct{}{}
	}

	var result []int
	for _, elem := range set2 {
		if _, ok := m[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}
