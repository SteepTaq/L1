package main

import "fmt"

func main() {
	nums := []int{2, 22, 33, 77, 50, 67, 70, 80, 90, 100}
	in := make(chan int)
	out := make(chan int)

	go func() {
		defer close(in)
		for _, num := range nums {
			in <- num
		}
	}()

	go func() {
		defer close(out)
		for num := range in {
			result := num * 2
			out <- result
		}
	}()

	for result := range out {
		fmt.Printf("%d\n", result)
	}
}
