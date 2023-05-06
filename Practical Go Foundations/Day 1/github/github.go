package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/AlexTLDR")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}
}
