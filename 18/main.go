package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Incr() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	counter := &Counter{}
	numWorkers := 6
	var wg sync.WaitGroup

	fmt.Printf("запустил %d горутин каждая выполнит %d инкрементов\n",
		numWorkers, 1000)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter.GetValue()) // ждём 6000

}
