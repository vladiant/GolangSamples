package main

// WorkerPools in Go Tutorial
// https://www.youtube.com/watch?v=1iBj5qVyfQA

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d  URL: %s\n", wId, site.URL)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{
			Status: resp.StatusCode,
			URL:    site.URL,
		}
	}
}

func main() {
	fmt.Println("worker pools in Go")

	const WorkerPoolsSize = 3

	jobs := make(chan Site, WorkerPoolsSize)
	results := make(chan Result, WorkerPoolsSize)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com",
		"https://google.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for range urls {
		result := <-results
		log.Println(result)
	}

}
