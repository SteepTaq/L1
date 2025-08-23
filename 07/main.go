package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	data map[string]int
	mu   sync.Mutex
}

func main() {
	fmt.Println("1)Mutex:")
	testMutex()

	fmt.Println("\n2)sync.Map:")
	testSyncMap()

	fmt.Println("\nвсё завершено")
}

func testMutex() {
	concurrentMap := &ConcurrentMap{
		data: make(map[string]int),
	}
	var wg sync.WaitGroup
	numGoroutines := 10
	iter := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iter; j++ {
				key := fmt.Sprintf("mutex_%d_key_%d", id, j)
				concurrentMap.Put(key, id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("  Записано %d элементов\n", len(concurrentMap.data))
}

func testSyncMap() {
	var syncMap sync.Map
	var wg sync.WaitGroup
	numWorkers := 10
	iter := 100

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iter; j++ {
				key := fmt.Sprintf("syncmap_%d_key_%d", id, j)
				syncMap.Store(key, id)
			}
		}(i)
	}

	wg.Wait()

	count := 0
	syncMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("  Записано %d элементов\n", count)
}

func (sm *ConcurrentMap) Put(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}
