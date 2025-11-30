package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//noWait()
	//withWait()
	//writeWithoutConcurrency()
	//writeWithConcurrency()
	//writeWithMutex()
	baseSelect()
}

func noWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(i + 1)
		}(i)
	}
	wg.Wait()
	fmt.Println("exit")
}

func writeWithoutConcurrency() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond) // simulate some work
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithConcurrency() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond) // simulate some work
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond) // simulate some work

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func baseSelect() {
	bufferedChannel := make(chan string, 1)
	bufferedChannel <- "first"
	select {
	//case str := <-bufferedChannel:
	//	fmt.Println("read", str)
	case bufferedChannel <- "second":
		fmt.Println("write", <-bufferedChannel, <-bufferedChannel)
	default:
		fmt.Println("default")
	}
}
