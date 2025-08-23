package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("1. Выход по условию:")
	stopByCondition()

	fmt.Println("\n2. Остановка через канал:")
	stopByChannel()

	fmt.Println("\n3. Остановка через контекст:")
	stopByContext()

	fmt.Println("\n4. Остановка через runtime.Goexit():")
	stopByGoexit()

	fmt.Println("\n5. Остановка через sync.WaitGroup:")
	stopByWaitGroup()

	fmt.Println("\n6. Остановка через select с таймером:")
	stopByTimer()

	fmt.Println("\n7. Остановка через panic/recover:")
	stopByPanic()

	fmt.Println("\n8. Остановка через сигналы:")
	stopBySignal()

	fmt.Println("\n9. Остановка через atomic:")
	stopByAtomic()
	fmt.Println("\nВсе примеры завершены")
}

// 1)Выход по условию
func stopByCondition() {
	var wg sync.WaitGroup
	stop := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if stop {
				fmt.Println("  Горутина остановлена по условию")
				return
			}
			fmt.Printf("  Итерация %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("  Горутина завершилась естественным образом")
	}()

	time.Sleep(300 * time.Millisecond)
	stop = true
	wg.Wait()
}

// 2)Остановка через канал
func stopByChannel() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				fmt.Println("  Горутина остановлена через канал")
				return
			default:
				fmt.Println("  Горутина работает")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	close(stopCh)
	wg.Wait()
}

// 3)Остановка через контекст
func stopByContext() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("  Горутина остановлена через контекст:", ctx.Err())
				return
			default:
				fmt.Println("  Горутина работает с контекстом")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	cancel()
	wg.Wait()
}

// 4)Остановка runtime.Goexit()
func stopByGoexit() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			fmt.Println("  Горутина завершается через defer")
		}()

		go func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("  Вызываем runtime.Goexit()")
			runtime.Goexit()
		}()

		for i := 0; i < 10; i++ {
			fmt.Printf("  Внутренняя горутина: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}

// 5)Остановка через sync.WaitGroup
func stopByWaitGroup() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stopCh:
					fmt.Printf("  Горутина %d остановлена\n", id)
					return
				default:
					fmt.Printf("  Горутина %d работает\n", id)
					time.Sleep(150 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(400 * time.Millisecond)
	close(stopCh)
	wg.Wait()
	fmt.Println("  Все горутины остановлены")
}

// 6)Остановка через select с таймером
func stopByTimer() {
	var wg sync.WaitGroup
	timer := time.NewTimer(500 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("  Горутина остановлена по таймеру")
				return
			default:
				fmt.Println("  Горутина работает с таймером")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

// 7)Остановка panic/recover
func stopByPanic() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Горутина восстановлена после panic: %v\n", r)
			}
		}()

		for i := 0; i < 5; i++ {
			fmt.Printf("  Итерация %d\n", i)
			time.Sleep(100 * time.Millisecond)

			if i == 3 {
				fmt.Println("  Вызываем panic")
				panic("Тест panic")
			}
		}
	}()

	wg.Wait()
}

// 8) Остановка сигналом
func stopBySignal() {
	var wg sync.WaitGroup
	signalCh := make(chan os.Signal, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case sig := <-signalCh:
				fmt.Printf("  Горутина остановлена сигналом: %v\n", sig)
				return
			default:
				fmt.Println("  Горутина ожидает сигнал")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(400 * time.Millisecond)
	signalCh <- os.Interrupt
	wg.Wait()
}

// 9) Остановка через atomic
func stopByAtomic() {
	var wg sync.WaitGroup
	var stopFlag int32

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if atomic.LoadInt32(&stopFlag) == 1 {
				fmt.Println("  Горутина остановлена через atomic")
				return
			}
			fmt.Println("Горутина работает")
			time.Sleep(200 * time.Millisecond)
		}
	}()
	time.Sleep(400 * time.Millisecond)
	atomic.StoreInt32(&stopFlag, 1)
	wg.Wait()
}
