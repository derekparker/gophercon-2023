package count

import (
	"io"
	"os"
	"strings"
	"sync"
)

// Job is a single job to be performed.
type Job struct {
	url string
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
		jobs = append(jobs, &Job{url: url})
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
