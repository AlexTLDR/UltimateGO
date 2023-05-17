package main

import (
	"fmt"
	"sync"
	"time"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // Dollar

	once sync.Once
}

func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() { p.process(t) })
}
func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s\n", ts, p.From, p.Amount, p.To)
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
