package main

import (
	"io"
	"log"
	"net/http"
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
}
