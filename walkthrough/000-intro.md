# Introduction

## Introduce first program

Run the following command:

```
go mod init gc2023
```

Then introduce the initial program:

```
package main

import (
"bufio"
"fmt"
"io"
"net/http"
"os"
"strings"
"time"
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

func worker(id int, urls <-chan string, results chan<- int) {
    for url := range urls {
        results <- doWork(id, url)
    }
}

func main() {
    var (
    finalResult int
    numJobs     int

    workers = 5
    urls    = make(chan string)
    results = make(chan int)
    )

    scanner := bufio.NewScanner(os.Stdin)
    start := time.Now()
    for w := 1; w <= workers; w++ {
        go worker(w, urls, results)
    }
    for scanner.Scan() {
        urls <- scanner.Text()
        numJobs++
    }
    close(urls)
    for num := range results {
        finalResult += num
    }
    t := time.Now()
    elapsed := t.Sub(start)
    for i := 1; i <= numJobs; i++ {
        oneResult := <-results
        finalResult += oneResult
    }
    fmt.Println("Number = ", finalResult)
    fmt.Println("Time = ", elapsed)
    if err := scanner.Err(); err != nil {
        _, _ = fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}

```

Then create a Makefile:

```
touch Makefile
```

And insert the following:

```
.PHONY: run
run:
    go run main.go < urls.txt
```

Finally, create a text file:

```
touch urls.txt
```

And insert the following:

```
http://go.dev
http://google.com
http://yahoo.com
http://wikipedia.com
http://bing.com
```

[Next](001.md)
