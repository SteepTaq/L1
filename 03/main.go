package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range jobs {
		fmt.Printf("worker %d: %d\n", id, v)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "write num workers\n")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers < 1 {
		fmt.Fprintln(os.Stderr, "num_workers должен быть положительным целым числом")
		os.Exit(1)
	}

	jobs := make(chan int, numWorkers*2)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	for value := 0; ; value++ {
		jobs <- value
		time.Sleep(100 * time.Millisecond)
	}
}
