package main

import "fmt"

func main() {
	nums := []int{64, 3, 25, 228, 76, 455, 33, 67, 89, 14, 89}

	fmt.Printf("исходный массив: %v\n", nums)
	sorted := quickSort(nums)
	fmt.Printf("отсортированный массив: %v\n", sorted)

}

func quickSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	opr := ar[len(ar)/2]
	var left, mid, right []int

	for _, num := range ar {
		switch {
		case num < opr:
			left = append(left, num)
		case num == opr:
			mid = append(mid, num)
		default:
			right = append(right, num)
		}
	}
	return append(append(quickSort(left), mid...), quickSort(right)...)
}
