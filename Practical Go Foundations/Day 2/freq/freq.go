package main

import (
	"io"
	"log"
	"os"
)

// What is the most common word in sherlock.txt ?

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer file.Close()

	wordFrequency(file)
}

func wordFrequency(r io.Reader) (map[string]int, error) {

}
