package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println(binSearch(nums, 15))
	fmt.Println(binSearch(nums, 5))
	fmt.Println(binSearch(nums, 10)) // -1
}
func binSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
