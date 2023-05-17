package main

import (
	"fmt"
	"sync"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // Dollar

	once sync.Once
}

func (p *Payment) Process() {
	p.once.Do(p.process)
}
func (p *Payment) process() {
	fmt.Printf("%s -> $%.2f -> %s\n", p.From, p.Amount, p.To)
}

func main() {
	p := Payment{
		From:   "Alex",
		To:     "Mazda",
		Amount: 99.99,
	}

	p.Process()
	p.Process()
}
