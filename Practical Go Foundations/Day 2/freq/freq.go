package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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
	mapDemo()
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	lnum := 0
	for s.Scan() {
		lnum++
		words := wordRe.FindAllString(s.Text(), -1) // current line
		if len(words) != 0 {
			fmt.Println(words)
			break
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Println("number of lines:", lnum)
	return nil, nil
}

func mapDemo() {
	var stocks map[string]float64 // word -> count
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}
	stocks = make(map[string]float64)
	stocks[sym] = 136.73
	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}
}
