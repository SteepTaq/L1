package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("до: %v\n", nums)
	nums1 := deleteNum(nums, 2)
	fmt.Printf("после: %v\n", nums1)
}

func deleteNum(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}
