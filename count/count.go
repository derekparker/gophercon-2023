package count

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Job is a single job to be performed.
type Job struct {
	Url string
}

// ParseURLsFromFile takes a path to a text file containing
// newline delimited URLs and parses them into Job structs
// which can be passed to workers.
func ParseURLsFromFile(path string) ([]*Job, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	list := string(b)
	urls := strings.Split(list, "\n")

	jobs := make([]*Job, len(urls))
	for _, url := range urls {
		jobs = append(jobs, &Job{Url: url})
	}

	return jobs, nil
}

func Worker(wg *sync.WaitGroup, id int, urls <-chan string, results chan<- int) {
	wg.Add(1)
	defer wg.Done()
	for url := range urls {
		results <- doWork(id, url)
	}
}

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
