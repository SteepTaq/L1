package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("Последовательность: %v\n", sequence)

	uniqSet := createSet(sequence)
	fmt.Printf("Множество = %v\n", uniqSet)
}

func createSet(elems []string) []string {
	m := make(map[string]struct{})

	for _, elem := range elems {
		m[elem] = struct{}{}
	}
	result := make([]string, 0, len(m))
	for elem := range m {
		result = append(result, elem)
	}

	return result
}
