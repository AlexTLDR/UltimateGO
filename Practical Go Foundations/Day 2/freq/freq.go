package main

import (
	"bufio"
	"fmt"
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
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		lnum++
		s.Text() // current line

	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("number of lines:", lnum)
	return nil, nil
}
