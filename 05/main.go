package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "write n sec\n")
		os.Exit(1)
	}

	sec, err := strconv.Atoi(os.Args[1])
	if err != nil || sec < 1 {
		fmt.Fprintln(os.Stderr, "количество секунд должно быть положительным целым числом")
		os.Exit(1)
	}

	fmt.Printf("программа работает %d секунд\n", sec)

	res := make(chan int)

	go func() {
		for value := 0; ; value++ {
			res <- value
			fmt.Printf("отправлено: %d\n", value)
			time.Sleep(222 * time.Millisecond)
		}
	}()

	timer := time.NewTimer(time.Duration(sec) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Printf("\nвремя истекло (%d секунд). завершение программы.\n", sec)
			return
		case value := <-res:
			fmt.Printf("прочитано: %d\n", value)
		}
	}
}
