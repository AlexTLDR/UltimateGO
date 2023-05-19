package main

import "fmt"

type filter interface {
	isInteresting(int) bool
}

type lessThan struct {
	baseLine int
}

type isOdd struct {
}

func (l lessThan) isInteresting(in int) bool {
	return l.baseLine < in
}

func (i isOdd) isInteresting(in int) bool {
	return in%2 != 0
}

func main() {
	numbers := []int{1, 2, 3, 4, 11, 12, 13}

	funcs := []filter{lessThan{baseLine: 10}, isOdd{}}
	interestingNumbers := []int{}
	for _, n := range numbers {
		isNumberInteresting := true
		for _, f := range funcs {
			isNumberInteresting = isNumberInteresting && f.isInteresting(n)
		}
		if isNumberInteresting {
			interestingNumbers = append(interestingNumbers, n)
		}
	}
	fmt.Println(interestingNumbers)
}
