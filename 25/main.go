package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	done := make(chan bool)

	go func() {
		timer := time.NewTimer(duration)
		<-timer.C
		done <- true
	}()
	<-done
}

func SleepLoop(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}

func main() {
	fmt.Println("начало Sleep")
	start := time.Now()
	Sleep(3 * time.Second)
	fmt.Printf("время истекло: %v\n", time.Since(start))

	fmt.Println("начало SleepLoop")
	start = time.Now()
	SleepLoop(3 * time.Second)
	fmt.Printf("время истекло: %v\n", time.Since(start))

}
