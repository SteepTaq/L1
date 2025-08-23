package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println("Исходный массив:", numbers)
	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			results <- square
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\nРезультаты:")

	var squares []int
	for square := range results {
		squares = append(squares, square)
	}

	for i, num := range numbers {
		fmt.Printf("%d**2 = %d\n", num, squares[i])
	}

	fmt.Println("\nВсе квадраты:", squares)
}
