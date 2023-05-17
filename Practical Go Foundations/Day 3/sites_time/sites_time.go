package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
		return
	}

	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}

func main() {
	//siteTime("https://google.com")
	urls := []string{
		"https://google.com",
		"https://gsp.ro",
		"https://amazon.de",
		"https://fake-site.org",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		// wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	wg.Wait()
}
