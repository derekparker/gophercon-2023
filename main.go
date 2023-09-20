package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/derekparker/gophercon-2023/counter"
)

func doWork(id int, url string) int {
	var data string
	fmt.Println("worker", id, "started  job", url)
	if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "https") {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "unable to close response body: %v", err)
			}
		}(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		data = string(body)
	} else {
		body, err := os.ReadFile(url)
		if err != nil {
			fmt.Println(err)
		}
		data = string(body)
	}
	number := strings.Count(data, "Go")
	fmt.Println("worker", id, "finished  job", url, "Number of Go is", number)
	return number
}

func main() {
	var (
		finalResult int
		wg          sync.WaitGroup

		workers = 5
		urls    = make(chan string)
		results = make(chan int, 5)
	)

	scanner := bufio.NewScanner(os.Stdin)
	start := time.Now()
	for w := 1; w <= workers; w++ {
		go counter.Worker(&wg, w, urls, results)
	}
	for scanner.Scan() {
		urls <- scanner.Text()
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
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
