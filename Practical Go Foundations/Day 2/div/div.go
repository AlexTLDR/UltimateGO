package main

import "fmt"

func main() {
	//fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
}

func safeDiv(a, b int) (int, error) {

	defer func() {
		if e := recover(); e != nil {

		}
	}()
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
