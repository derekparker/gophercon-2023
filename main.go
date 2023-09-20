package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/derekparker/gophercon-2023/count"
)

func main() {
	var (
		wg sync.WaitGroup

		finalResult int
		path        string

		workers = 5
		urls    = make(chan string)
		results = make(chan int, 5)
	)

	flag.StringVar(&path, "path", "", "text file containing list of URLs to scan")
	flag.Parse()

	jobs, err := count.ParseURLsFromFile(path)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	start := time.Now()
	for w := 1; w <= workers; w++ {
		go count.Worker(&wg, w, urls, results)
	}
	for _, job := range jobs {
		urls <- job.Url
	}
	close(urls)

	wg.Wait()
	close(results)

	for num := range results {
		finalResult += num
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Number = ", finalResult)
	fmt.Println("Time = ", elapsed)
}
