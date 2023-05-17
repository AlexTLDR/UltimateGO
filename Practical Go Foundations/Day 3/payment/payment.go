package main

import "fmt"

type Payment struct {
	From   string
	To     string
	Amount float64 // Dollar
}

func (p *Payment) Process() {
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
