package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: завершение работы\n", id)
			return
		case v, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("worker %d: %d\n", id, v)
		}
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		fmt.Printf("\nПолучен сигнал %v\n", sig)
		cancel()
	}()

	jobs := make(chan int, numWorkers*2)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	for value := 0; ; value++ {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return
		case jobs <- value:
		}
		time.Sleep(100 * time.Millisecond)
	}
}
