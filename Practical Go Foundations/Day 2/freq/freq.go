package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// What is the most common word (ignoring case) in sherlock.txt ?

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer file.Close()

	wordFrequency(file)
	// mapDemo()
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""

	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word count

	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}

// func mapDemo() {
// 	var stocks map[string]float64
// 	sym := "TTWO"
// 	price := stocks[sym]
// 	fmt.Printf("%s -> $%.2f\n", sym, price)

// 	if price, ok := stocks[sym]; ok {
// 		fmt.Printf("%s -> $%.2f\n", sym, price)
// 	} else {
// 		fmt.Printf("%s not found\n", sym)
// 	}
// 	stocks = make(map[string]float64)
// 	stocks[sym] = 136.73
// 	if price, ok := stocks[sym]; ok {
// 		fmt.Printf("%s -> $%.2f\n", sym, price)
// 	} else {
// 		fmt.Printf("%s not found\n", sym)
// 	}

// 	delete(stocks, "AAPL")
// }
