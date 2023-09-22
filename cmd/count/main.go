package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/derekparker/gophercon-2023/count"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	urls, ok := r.URL.Query()["urls"]
	if !ok || len(urls[0]) < 1 {
		log.Println("Url Param 'urls' is missing")
		return
	}

	var (
		wg          sync.WaitGroup
		finalResult int

		workers  = 5
		urlsChan = make(chan string, len(urls))
		results  = make(chan int, len(urls))
		start    = time.Now()
	)

	for _, url := range urls {
		urlsChan <- url
	}
	close(urlsChan)

	for w := 1; w <= workers; w++ {
		go count.Worker(&wg, w, urlsChan, results)
	}

	wg.Wait()

	for num := range results {
		finalResult += num
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Fprintln(w, "Number = ", finalResult)
	fmt.Fprintln(w, "Time = ", elapsed)
}

func main() {
	http.HandleFunc("/count", countHandler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
