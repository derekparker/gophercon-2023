package main

import (
	"fmt"
	"sync"
	"time"
)

const workers = 5

func doWork(wg *sync.WaitGroup, id int) {
	fmt.Println("worker", id, "starting...")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("worker", id, "finishing...")

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go doWork(&wg, i)
	}

	wg.Wait()
}
